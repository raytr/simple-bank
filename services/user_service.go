package services

import (
	"context"

	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/helper/password"
	"gibhub.com/raytr/simple-bank/helper/random"
	"gibhub.com/raytr/simple-bank/models/entity"
	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/repository"
	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

type UserService interface {
	Save(ctx context.Context, input request.UserRegister) error
	GetById(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetList(ctx context.Context, input request.ListUserRequest) ([]entity.User, error)
	Update(ctx context.Context, id uuid.UUID, input request.UpdateUserBody) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userServiceImpl struct {
	repository repository.UserRepo
	secConfig  config.SecurityConfig
	logger     log.Logger
}

func NewUserService(
	userRepo repository.UserRepo,
	secCfg config.SecurityConfig,
	logger log.Logger,
) UserService {
	return &userServiceImpl{userRepo, secCfg, logger}
}

func (s *userServiceImpl) Save(ctx context.Context, input request.UserRegister) error {
	salt := random.GenerateRandomString(s.secConfig.PasswordSaltLength)
	hashedPassword, err := password.HashPassword(input.Password, salt, s.secConfig.PasswordPepper)
	if err != nil {
		return err
	}

	user := &entity.User{
		FullName:     input.FullName,
		Username:     input.Username,
		HashPassword: hashedPassword,
		Salt:         salt,
	}

	err = s.repository.Create(ctx, user)
	if err != nil {
		s.logger.Log(err)
		return err
	}

	return nil
}

func (s *userServiceImpl) GetById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := s.repository.FindById(ctx, id)
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) GetList(ctx context.Context, filter request.ListUserRequest) ([]entity.User, error) {
	users, err := s.repository.FindByParams(ctx, filter)
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}

	return users, nil
}

func (s *userServiceImpl) Update(ctx context.Context, id uuid.UUID, input request.UpdateUserBody) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		s.logger.Log(err)
		return err
	}

	data := make(map[string]interface{})
	if input.FullName != nil {
		data["full_name"] = *input.FullName
	}

	err = s.repository.Update(ctx, id, data)
	if err != nil {
		s.logger.Log(err)
		return err
	}

	return nil
}

func (s *userServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		s.logger.Log(err)
		return err
	}

	err = s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log(err)
		return err
	}

	return nil
}
