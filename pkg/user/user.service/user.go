package userservice

import (
	"errors"
	"first-project/db/model/domain"
	"first-project/db/model/entity"
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

func (uS *UserService) GetID(id int) (entity.UserHistoryEntity, error) {
	// tokenV, errToken := uS.Token.DecodeToken(token)
	// if errToken != nil {
	// 	return entity.UserHistoryEntity{}, errToken
	// }

	// claims, _ := tokenV.Claims.(*helper.CustomClaims)
	user, errUser := uS.Repo.GetID(id)

	if errUser != nil {
		return entity.UserHistoryEntity{}, nil
	}

	return entity.ToUserHistoryEntity(user), nil
}

func (uS *UserService) Update(id int, req web.UserUpdateReq) (helper.CustomResponse, error) {
	userId, errId := uS.Repo.GetID(id)

	if errId != nil {
		return nil, errId
	}

	// hash password if req is not empty
	var hashPass string
	if req.Password != "" {
		hashPassByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		hashPass = string(hashPassByte)
	} else {
		hashPass = userId.Password
	}

	req.Name = helper.DefaultEmpty(req.Name, userId.Name).(string)
	req.Email = helper.DefaultEmpty(req.Email, userId.Email).(string)
	req.Password = hashPass

	dataUpdate := domain.User{
		UserID:   id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	updateUser, errUpdate := uS.Repo.Update(dataUpdate)

	if errUpdate != nil {
		return nil, errUpdate
	}

	data := helper.CustomResponse{
		"name":     updateUser.Name,
		"email":    updateUser.Email,
		"password": updateUser.Password,
	}

	return data, nil
}
