package userrepository

import "first-project/db/model/domain"

type UserRepoInterface interface {
	Create(user domain.User) (domain.User, error)
	GetEmail(email string) (domain.User, error)
	GetID(id int) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id int) error
}
