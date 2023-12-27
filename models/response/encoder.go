package response

import (
	"context"
	"encoding/json"
	"errors"
	"gibhub.com/raytr/simple-bank/helper/b_error"
	"net/http"
)

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, b_error.ErrBadRequest):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, b_error.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

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
