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