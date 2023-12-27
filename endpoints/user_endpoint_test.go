package endpoints

import (
	"context"
	"gibhub.com/raytr/simple-bank/models/entity"
	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/services/services_mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeGetUsers(t *testing.T) {
	userSvcMock := services_mock.NewUserService(t)
	ctx := context.Background()

	entityUser := entity.User{
		BaseModel: entity.BaseModel{
			Id: uuid.New(),
		},
		FullName:     "full name",
		Username:     "username",
		HashPassword: "hash password",
	}

	filter := request.ListUserRequest{}

	userSvcMock.Mock.On("GetList", ctx, filter).Return([]entity.User{entityUser}, nil)
	endpoint := makeGetUsers(userSvcMock)

	t.Run("success", func(t *testing.T) {
		_, err := endpoint(ctx, filter)
		require.NoErrorf(t, err, "expected %v received %v", nil, err)
	})
}

func TestMakeSaveUser(t *testing.T) {
	userSvcMock := services_mock.NewUserService(t)
	ctx := context.Background()

	req := request.UserRegister{
		FullName: "full name",
		Username: "username",
		Password: "password",
	}

	userSvcMock.Mock.On("Save", ctx, req).Return(nil)
	endpoint := makeSaveUser(userSvcMock)

	t.Run("success", func(t *testing.T) {
		_, err := endpoint(ctx, req)
		require.NoErrorf(t, err, "expected %v received %v", nil, err)
	})
}

func TestMakeUpdateUser(t *testing.T) {
	userSvcMock := services_mock.NewUserService(t)
	ctx := context.Background()

	req := request.UpdateUserRequest{
		ID:   uuid.New(),
		Body: request.UpdateUserBody{},
	}

	userSvcMock.Mock.On("Update", ctx, req.ID, req.Body).Return(nil)
	endpoint := makeUpdateUser(userSvcMock)

	t.Run("success", func(t *testing.T) {
		_, err := endpoint(ctx, req)
		require.NoErrorf(t, err, "expected %v received %v", nil, err)
	})
}

func TestMakeDeleteUser(t *testing.T) {
	userSvcMock := services_mock.NewUserService(t)
	ctx := context.Background()

	id := uuid.New()

	userSvcMock.Mock.On("Delete", ctx, id).Return(nil)
	endpoint := makeDeleteUser(userSvcMock)

	t.Run("success", func(t *testing.T) {
		_, err := endpoint(ctx, id)
		require.NoErrorf(t, err, "expected %v received %v", nil, err)
	})
}
