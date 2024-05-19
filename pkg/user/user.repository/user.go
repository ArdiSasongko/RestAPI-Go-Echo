package userrepository

import (
	"errors"
	"first-project/db/model/domain"
	"strings"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (uR *UserRepo) Create(user domain.User) (domain.User, error) {
	if err := uR.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.User{}, errors.New("email already exists")
		}
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}

func (uR *UserRepo) GetEmail(email string) (domain.User, error) {
	var user domain.User
	if err := uR.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		return domain.User{}, errors.New("email not found")
	}
	return user, nil
}

func (uR *UserRepo) GetID(id int) (domain.User, error) {
	var user domain.User
	if err := uR.DB.Preload("Orders").First(&user, "user_id = ?", id).Error; err != nil {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (uR *UserRepo) Update(user domain.User) (domain.User, error) {
	if err := uR.DB.Model(domain.User{}).Where("user_id = ?", user.UserID).Updates(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return domain.User{}, errors.New("email already exists")
		}
		return domain.User{}, err
	}

	return user, nil
}
