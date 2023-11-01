package endpoints

import (
	"context"

	"gibhub.com/raytr/simple-bank/models/request"
	"gibhub.com/raytr/simple-bank/models/response"
	"gibhub.com/raytr/simple-bank/services"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type UserEndpoint struct {
	List     endpoint.Endpoint
	Register endpoint.Endpoint
	Update   endpoint.Endpoint
	Delete   endpoint.Endpoint
}

func MakeUserEndpoints(userSvc services.UserService) UserEndpoint {
	return UserEndpoint{
		List:     makeGetUsers(userSvc),
		Register: makeSaveUser(userSvc),
		Update:   makeUpdateUser(userSvc),
		Delete:   makeDeleteUser(userSvc),
	}
}

func makeGetUsers(userSvc services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.ListUserRequest)

		result, err := userSvc.GetList(ctx, req)
		if err != nil {
			return nil, err
		}

		collection := response.UserRegisterCollection(result)
		return collection.ToResponse(), nil
	}
}

func makeSaveUser(userSvc services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.UserRegister)
		err = userSvc.Save(ctx, req)
		if err != nil {
			return nil, err
		}

		return response.SuccessResponse, nil
	}
}

func makeUpdateUser(userSvc services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.UpdateUserRequest)
		err = userSvc.Update(ctx, req.ID, req.Body)
		if err != nil {
			return nil, err
		}

		return response.SuccessResponse, nil
	}
}

func makeDeleteUser(userSvc services.UserService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		id := rqst.(uuid.UUID)
		err = userSvc.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		return response.SuccessResponse, nil
	}
}
