package response

import "gibhub.com/raytr/simple-bank/models/entity"

type UserRegisterResponse struct {
	Username string `json:"username"`
}

type UserRegisterCollection []entity.User

// ToResponse func will serialize model entity slice to response entity slice
func (collection UserRegisterCollection) ToResponse() []UserRegisterResponse {
	a := len(collection)
	response := make([]UserRegisterResponse, 0, a)
	for _, user := range collection {
		ur := UserRegisterResponse{
			Username: user.Username,
		}
		response = append(response, ur)
	}

	return response
}
