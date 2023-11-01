package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"gibhub.com/raytr/simple-bank/endpoints"
	"gibhub.com/raytr/simple-bank/helper/password"
	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/services"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func AuthHttpHandler(authSvc services.AuthService, logger kitlog.Logger) *mux.Router {
	router := mux.NewRouter()
	epAuth := endpoints.MakeAuthEndpoints(authSvc)
	RegisterAuthHttpHandler(router, epAuth, logger)

	return router
}

func RegisterAuthHttpHandler(r *mux.Router, ep endpoints.AuthEndpoint, logger kitlog.Logger) {

	//options provided by the Go kit to facilitate error control
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(password.EncodeErrorResponse),
	}

	r.Methods("POST").Path("/login").Handler(httptransport.NewServer(
		ep.Login,
		decodeRegisterAuth,
		password.EncodeResponse,
		options...,
	))

	r.Methods("POST").Path("/renew-access-token").Handler(httptransport.NewServer(
		ep.RenewAccessToken,
		decodeRenewAccessToken,
		password.EncodeResponse,
		options...,
	))

}

func decodeRegisterAuth(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var req request.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeRenewAccessToken(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var req request.RenewAccessTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
