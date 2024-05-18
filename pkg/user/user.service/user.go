package userservice

import (
	"errors"
	"first-project/db/model/domain"
	"first-project/db/model/web"
	"first-project/helper"
	userrepository "first-project/pkg/user/user.repository"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo  userrepository.UserRepoInterface
	Token helper.TokenUseCaseInterface
}

func NewUserService(repo userrepository.UserRepoInterface, token helper.TokenUseCaseInterface) *UserService {
	return &UserService{
		Repo:  repo,
		Token: token,
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

func (uS *UserService) Login(email, password string) (helper.CustomResponse, error) {
	user, errUser := uS.Repo.GetEmail(email)

	if errUser != nil {
		return nil, errUser
	}

	if errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); errPass != nil {
		return nil, errors.New("password invalid")
	}

	expiredTime := time.Now().Local().Add(5 * time.Minute)

	claims := helper.CustomClaims{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Echo",
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, errToken := uS.Token.GeneratedToken(claims)

	if errToken != nil {
		return nil, errToken
	}

	data := helper.CustomResponse{
		"token":     token,
		"expiredAt": expiredTime,
	}

	return data, nil
}
