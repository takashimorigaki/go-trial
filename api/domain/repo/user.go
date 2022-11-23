package repo

import "go-trial/domain/model"

type UserRepo interface {
	GetList() ([]model.User, error)
	Create(user model.User) (model.User, error)
}