package repoimple

import (
	"go-trial/domain/model"

	"github.com/jinzhu/gorm"
)

type UserRepoimple struct {
	db *gorm.DB
}

func SetUserRepoimple(db *gorm.DB) UserRepoimple {
	return UserRepoimple{db: db}
}

func (ur UserRepoimple) GetList() ([]model.User, error) {
	db := ur.db

	var	users []model.User
	err := db.Order("created_at asc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur UserRepoimple) GetSingle(id string) (model.User, error) {
	db := ur.db

	var user model.User
	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur UserRepoimple) Create(user model.User) (model.User, error) {
	db := ur.db

	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur UserRepoimple) Update(user model.User) (model.User, error) {
	db := ur.db

	err := db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}