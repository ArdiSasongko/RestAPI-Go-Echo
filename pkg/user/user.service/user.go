package userservice

import (
	"first-project/db/model/domain"
	"first-project/db/model/web"
	"first-project/helper"
	userrepository "first-project/pkg/user/user.repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo userrepository.UserRepoInterface
}

func NewUserService(repo userrepository.UserRepoInterface) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (uS *UserService) Create(req web.UserReq) (helper.CustomResponse, error) {
	passHash, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)

	if errHash != nil {
		return nil, errHash
	}

	newUser := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passHash),
	}

	saveUser, errSave := uS.Repo.Create(newUser)

	if errSave != nil {
		return nil, errSave
	}

	dataUser := helper.CustomResponse{
		"user_id":  saveUser.UserID,
		"name":     saveUser.Name,
		"email":    saveUser.Email,
		"password": saveUser.Password,
	}

	return dataUser, nil
}
