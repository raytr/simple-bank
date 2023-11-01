package request

import (
	"regexp"

	"github.com/google/uuid"
	validation "github.com/itgelo/ozzo-validation/v4"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type ListUserRequest struct {
	FullName *string `schema:"full_name" json:"full_name"`
	Username *string `schema:"username" json:"username"`
}

type DeleteUserRequest struct {
	ID uuid.UUID
}

type UpdateUserRequest struct {
	ID   uuid.UUID `json:"id"`
	Body UpdateUserBody
}

type UpdateUserBody struct {
	FullName *string `json:"full_name"`
}

func (req UserRegister) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required.Error("username is required")),
		validation.Field(&req.FullName, validation.Required.Error("full_name is required")),
		validation.Field(&req.Password, validation.Required.Error("password is required")),
		validation.Field(&req.Username, validation.Required.Error("username is invalid"), validation.Match(regexp.MustCompile("^[A-Za-z0-9]*$"))),
		validation.Field(&req.Password, validation.Required.Error("password is invalid"), validation.Match(regexp.MustCompile("^[A-Za-z0-9]*$"))),
	)
}
