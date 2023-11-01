package endpoints

import (
	"context"

	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/services"
	"github.com/go-kit/kit/endpoint"
)

type AuthEndpoint struct {
	Login            endpoint.Endpoint
	RenewAccessToken endpoint.Endpoint
}

func MakeAuthEndpoints(authSvc services.AuthService) AuthEndpoint {
	return AuthEndpoint{
		Login:            makeLogin(authSvc),
		RenewAccessToken: makeRenewAccessToken(authSvc),
	}
}

func makeLogin(authSvc services.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.LoginRequest)

		response, err := authSvc.Login(ctx, req.Username, req.Password)
		if err != nil {
			return nil, err
		}

		return response, nil
	}
}

func makeRenewAccessToken(authSvc services.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.RenewAccessTokenRequest)

		response, err := authSvc.RenewAccessToken(ctx, req.RefreshToken)
		if err != nil {
			return nil, err
		}

		return response, nil
	}
}
