package password

import (
	"context"
	"encoding/json"
	"net/http"

	"gibhub.com/raytr/simple-bank/helper/b_error"
	"golang.org/x/crypto/bcrypt"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	return json.NewEncoder(w).Encode(response)
}

func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case b_error.ErrInvalidUser:
		return http.StatusNotFound
	case b_error.ErrInvalidToken:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func HashPassword(password string, salt string, pepper string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt+pepper), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPassword(password string, hashedPassword string, salt string, pepper string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt+pepper))
}
