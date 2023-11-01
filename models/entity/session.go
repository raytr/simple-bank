package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `gorm:"primarykey"`
	UserID       uuid.UUID `gorm:"not null"`
	RefreshToken string    `gorm:"type:varchar;not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	IsBlocked    bool      `gorm:"not null; default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
