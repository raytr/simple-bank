package token

import (
	"time"

	"gibhub.com/raytr/simple-bank/helper/b_error"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID uuid.UUID, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration * time.Minute),
	}

	return payload, err
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return b_error.ErrExpiredToken
	}
	return nil
}
