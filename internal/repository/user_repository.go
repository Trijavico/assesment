package repository

import (
	"context"

	"github.com/Trijavico/ensolvers-assesment/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(context.Context, *model.User) error
	FindByID(context.Context, uint) (model.User, error)
	FindByEmail(context.Context, string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (userRepo *userRepository) FindByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User

	err := userRepo.db.WithContext(ctx).Find(&user, id).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (userRepo *userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := userRepo.db.WithContext(ctx).Where(&model.User{Email: email}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (userRepo *userRepository) Save(ctx context.Context, user *model.User) error {
	err := userRepo.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
