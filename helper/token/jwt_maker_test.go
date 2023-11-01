package token

import (
	"testing"
	"time"

	"gibhub.com/raytr/simple-bank/helper/b_error"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {

	userID := uuid.New()
	duration := time.Minute
	secretKey := "secret"

	maker := NewJWTMaker(secretKey)

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration * time.Minute)

	token, _, err := maker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {

	userID := uuid.New()
	secretKey := "secret"

	payload, err := NewPayload(userID, time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker := NewJWTMaker(secretKey)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, b_error.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
