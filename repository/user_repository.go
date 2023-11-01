package repository

import (
	"context"

	"gibhub.com/raytr/simple-bank/models/entity"
	"gibhub.com/raytr/simple-bank/models/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepo interface {
	Create(ctx context.Context, user *entity.User) error
	FindByParams(ctx context.Context, filter request.ListUserRequest) ([]entity.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindOneByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, id uuid.UUID, input map[string]interface{}) error
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("id=?", id).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOneByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("username=?", username).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByParams(ctx context.Context, filter request.ListUserRequest) ([]entity.User, error) {
	var users []entity.User
	query := r.db.WithContext(ctx)

	if filter.Username != nil {
		query.Where("username = ?", filter.Username)
	}
	if filter.FullName != nil {
		query.Where("full_name = ?", filter.FullName)
	}

	return users, query.Find(&users).Error
}

func (r *userRepository) Update(ctx context.Context, id uuid.UUID, data map[string]interface{}) error {
	return r.db.WithContext(ctx).Table("users").Where("id=?", id).Updates(&data).Error
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	user := new(entity.User)
	return r.db.WithContext(ctx).Where("id=?", id).Delete(&user).Error
}
