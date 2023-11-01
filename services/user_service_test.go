package services

import (
	"context"
	"testing"

	"gibhub.com/raytr/simple-bank/config"
	"gibhub.com/raytr/simple-bank/models/entity"
	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/repository/repository_mock"
	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	logger log.Logger
	pwCfg  = config.SecurityConfig{
		PasswordPepper:     "pepper",
		PasswordSaltLength: 10,
	}
)

func TestRegisterUser(t *testing.T) {
	req := request.UserRegister{
		Username: "mock username",
		Password: "mock password",
		FullName: "mock full name",
	}

	userRepo := repository_mock.NewUserRepo(t)
	userSvc := NewUserService(userRepo, pwCfg, logger)
	ctx := context.Background()
	userRepo.Mock.On("Create", ctx, mock.Anything).Return(nil)

	err := userSvc.Save(ctx, req)
	require.NoError(t, err)
}

func TestGetByID(t *testing.T) {
	userRepo := repository_mock.NewUserRepo(t)
	userSvc := NewUserService(userRepo, pwCfg, logger)
	ctx := context.Background()

	userEntity := &entity.User{
		Username:     "mock username",
		HashPassword: "mock hash password",
		FullName:     "mock full name",
	}

	userRepo.Mock.On("FindById", ctx, userEntity.Id).Return(userEntity, nil)

	_, err := userSvc.GetById(ctx, userEntity.Id)
	require.NoError(t, err)
}

func TestGetList(t *testing.T) {
	userRepo := repository_mock.NewUserRepo(t)
	userSvc := NewUserService(userRepo, pwCfg, logger)
	ctx := context.Background()

	filter := request.ListUserRequest{}

	userEntity := []entity.User{
		{
			Username:     "mock username",
			HashPassword: "mock hash password",
			FullName:     "mock full name",
		},
	}

	userRepo.Mock.On("FindByParams", ctx, filter).Return(userEntity, nil)

	_, err := userSvc.GetList(ctx, filter)
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	userID := uuid.New()
	userRepo := repository_mock.NewUserRepo(t)
	userSvc := NewUserService(userRepo, pwCfg, logger)
	ctx := context.Background()

	updatingInfo := request.UpdateUserBody{}
	data := make(map[string]interface{})
	userEntity := &entity.User{
		BaseModel: entity.BaseModel{Id: userID},
	}

	userRepo.Mock.On("FindById", ctx, userEntity.Id).Return(userEntity, nil)
	userRepo.Mock.On("Update", ctx, userEntity.Id, data).Return(nil)

	err := userSvc.Update(ctx, userEntity.Id, updatingInfo)
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	userID := uuid.New()
	userRepo := repository_mock.NewUserRepo(t)
	userSvc := NewUserService(userRepo, pwCfg, logger)
	ctx := context.Background()

	userEntity := &entity.User{
		BaseModel: entity.BaseModel{Id: userID},
	}

	userRepo.Mock.On("FindById", ctx, userEntity.Id).Return(userEntity, nil)
	userRepo.Mock.On("Delete", ctx, userEntity.Id).Return(nil)

	err := userSvc.Delete(ctx, userEntity.Id)
	require.NoError(t, err)
}
