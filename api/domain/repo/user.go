package repo

import "go-trial/domain/model"

type UserRepo interface {
	GetList() (*[]model.User, error)
	GetSingle(id string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) (*model.User, error)
}