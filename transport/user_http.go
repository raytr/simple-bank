package transport

import (
	"context"
	"encoding/json"
	"gibhub.com/raytr/simple-bank/helper/b_log"
	"gibhub.com/raytr/simple-bank/models/response"
	"net/http"
	"strconv"

	"gibhub.com/raytr/simple-bank/endpoints"
	"gibhub.com/raytr/simple-bank/helper/password"
	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/services"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func UserHttpHandler(userSvc services.UserService, logger b_log.Logger) *mux.Router {
	router := mux.NewRouter()
	epUser := endpoints.MakeUserEndpoints(userSvc)
	RegisterUserHttpHandler(router, epUser, logger)

	return router
}

func RegisterUserHttpHandler(r *mux.Router, ep endpoints.UserEndpoint, logger b_log.Logger) {

	//options provided by the Go kit to facilitate error control
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(b_log.TraceIdentifier),
		httptransport.ServerErrorHandler(logger),
		httptransport.ServerErrorEncoder(response.EncodeError),
	}

	r.Methods("GET").Path("/users").Handler(httptransport.NewServer(
		ep.List,
		decodeListUser,
		password.EncodeResponse,
		options...,
	))

	r.Methods("POST").Path("/user/register").Handler(httptransport.NewServer(
		ep.Register,
		decodeRegisterUser,
		password.EncodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/user/{id}").Handler(httptransport.NewServer(
		ep.Update,
		decodeUpdateUser,
		password.EncodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/user/{id}").Handler(httptransport.NewServer(
		ep.Delete,
		decodeDeleteUser,
		password.EncodeResponse,
		options...,
	))
}

func decodeListUser(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var params request.ListUserRequest

	if err = schema.NewDecoder().Decode(&params, r.URL.Query()); err != nil {
		return nil, err
	}

	return params, nil
}

func decodeDeleteUser(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func decodeRegisterUser(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var req request.UserRegister

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	//validate
	validationErr := req.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	return req, nil
}

func decodeUpdateUser(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var req request.UpdateUserRequest

	idStr := mux.Vars(r)["id"]
	req.ID, err = uuid.FromBytes([]byte(idStr))
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, err
	}

	return req, nil
}
