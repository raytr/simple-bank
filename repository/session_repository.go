package repository

import (
	"context"

	"gibhub.com/raytr/simple-bank/models/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

type SessionRepo interface {
	CreateWithTx(ctx context.Context, session *entity.Session, tx *gorm.DB) error
	FindById(ctx context.Context, id uuid.UUID) (*entity.Session, error)
	Delete(ctx context.Context, id uint64) error
}

func NewSessionRepository(db *gorm.DB) SessionRepo {
	return &sessionRepository{db}
}

func (r *sessionRepository) CreateWithTx(ctx context.Context, session *entity.Session, tx *gorm.DB) error {
	return tx.WithContext(ctx).Create(session).Error
}

func (r *sessionRepository) FindById(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
	var session entity.Session
	err := r.db.WithContext(ctx).
		Where("id=?", id).
		First(&session).Error
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *sessionRepository) Delete(ctx context.Context, id uint64) error {
	session := new(entity.Session)
	return r.db.WithContext(ctx).Where("id=?", id).Delete(&session).Error
}
