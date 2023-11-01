package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"gibhub.com/raytr/simple-bank/helper/b_error"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) Maker {
	return &JWTMaker{secretKey: secretKey}
}

func (maker *JWTMaker) CreateToken(userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))

	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	funcKey := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, b_error.ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, funcKey)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, b_error.ErrExpiredToken) {
			return nil, b_error.ErrExpiredToken
		}
		return nil, b_error.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, b_error.ErrInvalidToken
	}

	return payload, nil
}
