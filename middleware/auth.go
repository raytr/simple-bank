package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gibhub.com/raytr/simple-bank/helper/token"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
)

type Auth struct {
	Pepper     string
	TokenMaker token.Maker
	Skip       bool
}

func (a *Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if !a.Skip {
			authorizationHeader := r.Header.Get(authorizationHeaderKey)
			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header was not provided")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "invalid authorization header format",
				})
				return
			}

			authorizationType := strings.ToTitle(fields[0])
			if strings.ToLower(authorizationType) != authorizationTypeBearer {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": fmt.Sprintf("unsupported authorization type %s", authorizationType),
				})
				return
			}

			accessToken := fields[1]
			payload, err := a.TokenMaker.VerifyToken(accessToken)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			ctx := context.WithValue(r.Context(), "user", payload)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
