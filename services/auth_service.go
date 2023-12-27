package services

import (
	"context"
	"errors"
	"gibhub.com/raytr/simple-bank/helper/b_log"
	"time"

	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/helper/b_error"
	passwordHelper "gibhub.com/raytr/simple-bank/helper/password"
	tokenHelper "gibhub.com/raytr/simple-bank/helper/token"
	"gibhub.com/raytr/simple-bank/models/entity"
	"gibhub.com/raytr/simple-bank/models/response"
	"gibhub.com/raytr/simple-bank/repository"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (*response.LoginResponse, error)
	RenewAccessToken(ctx context.Context, refreshToken string) (*response.LoginResponse, error)
}

type authServiceImpl struct {
	BaseRepo       repository.BaseRepository
	UserRepository repository.UserRepo
	SessionRepo    repository.SessionRepo
	SecurityConfig config.SecurityConfig
	TokenMaker     tokenHelper.Maker
	logger         b_log.Logger
}

func NewAuthService(
	baseRepo repository.BaseRepository,
	userRepo repository.UserRepo,
	sessionRepo repository.SessionRepo,
	secCfg config.SecurityConfig,
	tokenMaker tokenHelper.Maker,
	logger b_log.Logger,
) AuthService {
	return &authServiceImpl{baseRepo, userRepo, sessionRepo, secCfg, tokenMaker, logger}
}

func (s *authServiceImpl) RenewAccessToken(ctx context.Context, refreshToken string) (*response.LoginResponse, error) {
	refreshTokenPayload, err := s.TokenMaker.VerifyToken(refreshToken)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	session, err := s.SessionRepo.FindById(ctx, refreshTokenPayload.ID)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	if session.IsBlocked {
		err = errors.New("blocked session")
		s.logger.Error(err)
		return nil, err
	}

	if session.RefreshToken != refreshToken {
		err = errors.New("mismatched session token")
		s.logger.Error(err)
		return nil, err
	}

	if session.UserID != refreshTokenPayload.UserID {
		err = errors.New("incorrect session user")
		s.logger.Error(err)
		return nil, err
	}

	if time.Now().After(session.ExpiresAt) {
		err = errors.New("expired session")
		s.logger.Error(err)
		return nil, err
	}

	//=============

	accessToken, accessTokenPayload, err := s.TokenMaker.CreateToken(refreshTokenPayload.UserID, s.SecurityConfig.AccessTokenDuration)
	if err != nil {
		return nil, errors.New("occurred error when create access token: " + err.Error())
	}

	res := &response.LoginResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
	}

	return res, nil

}

func (s *authServiceImpl) Login(ctx context.Context, username, password string) (*response.LoginResponse, error) {
	user, err := s.UserRepository.FindOneByUsername(ctx, username)
	if err != nil {
		s.logger.Error(err)
		return nil, errors.New("wrong username or password")
	}

	err = passwordHelper.CheckPassword(password, user.HashPassword, user.Salt, s.SecurityConfig.PasswordPepper)
	if err != nil {
		return nil, b_error.ErrUnauthorized
	}

	tx := s.BaseRepo.GetBegin()

	accessToken, accessTokenPayload, err := s.TokenMaker.CreateToken(user.Id, s.SecurityConfig.AccessTokenDuration)
	if err != nil {
		return nil, errors.New("occurred error when create access token: " + err.Error())
	}

	refreshToken, refreshTokenPayload, err := s.TokenMaker.CreateToken(user.Id, s.SecurityConfig.RefreshTokenDuration)
	if err != nil {
		return nil, errors.New("occurred error when create refresh token: " + err.Error())
	}

	session := &entity.Session{
		ID:           refreshTokenPayload.ID,
		UserID:       user.Id,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
		IsBlocked:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = s.SessionRepo.CreateWithTx(ctx, session, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	res := &response.LoginResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
	}

	return res, nil
}
