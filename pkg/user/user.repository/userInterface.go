package userrepository

import "first-project/db/model/domain"

type UserRepoInterface interface {
	Create(user domain.User) (domain.User, error)
}