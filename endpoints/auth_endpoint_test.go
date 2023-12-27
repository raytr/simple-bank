package endpoints

import (
	"context"
	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/models/response"
	"gibhub.com/raytr/simple-bank/services/services_mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeLogin(t *testing.T) {
	authSvcMock := services_mock.NewAuthService(t)
	ctx := context.Background()

	req := request.LoginRequest{
		Username: "username",
		Password: "password",
	}

	res := &response.LoginResponse{
		Username: "username",
	}

	authSvcMock.Mock.On("Login", ctx, req.Username, req.Password).Return(res, nil)
	endpoint := makeLogin(authSvcMock)

	t.Run("success", func(t *testing.T) {
		_, err := endpoint(ctx, req)
		require.NoErrorf(t, err, "expected %v received %v", nil, err)
	})
}

func TestMakeRenewAccessToken(t *testing.T) {
	authSvcMock := services_mock.NewAuthService(t)
	ctx := context.Background()

	req := request.RenewAccessTokenRequest{
		RefreshToken: "refresh token",
	}

	res := &response.LoginResponse{
		Username: "username",
	}

	authSvcMock.Mock.On("RenewAccessToken", ctx, req.RefreshToken).Return(res, nil)
	endpoint := makeRenewAccessToken(authSvcMock)

	t.Run("success", func(t *testing.T) {
		_, err := endpoint(ctx, req)
		require.NoErrorf(t, err, "expected %v received %v", nil, err)
	})
}
