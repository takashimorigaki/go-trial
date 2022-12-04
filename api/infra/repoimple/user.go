package repoimple

import (
	"fmt"
	"go-trial/domain/model"
	"go-trial/lib/errorhandle"

	"net/http"

	"github.com/jinzhu/gorm"
)

type UserRepoimple struct {
	db *gorm.DB
}

func SetUserRepoimple(db *gorm.DB) UserRepoimple {
	return UserRepoimple{db: db}
}

func (ur UserRepoimple) GetList() (*[]model.User, error) {
	db := ur.db

	users := &[]model.User{}
	err := db.Order("created_at asc").Find(&users).Error
	if err != nil {
		cErr := errorhandle.InitCustomError(http.StatusNotFound, err)
		wrappedErr := fmt.Errorf("UserRepoimple.GetList: %w", cErr)		
		return nil, wrappedErr
	}

	return users, nil
}

func (ur UserRepoimple) GetSingle(id string) (*model.User, error) {
	db := ur.db

	user := &model.User{}
	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		cErr := errorhandle.InitCustomError(http.StatusNotFound, err)
		wrappedErr := fmt.Errorf("UserRepoimple.GetSingle: %w", cErr)		
		return nil, wrappedErr
	}
	return user, nil
}

func (ur UserRepoimple) Create(user *model.User) (*model.User, error) {
	db := ur.db

	err := db.Create(&user).Error
	if err != nil {
		cErr := errorhandle.InitCustomError(http.StatusBadRequest, err)
		wrappedErr := fmt.Errorf("UserRepoimple.Create: %w", cErr)		
		return nil, wrappedErr
	}
	return user, nil
}

func (ur UserRepoimple) Update(user *model.User) (*model.User, error) {
	db := ur.db

	err := db.Save(&user).Error
	if err != nil {
		cErr := errorhandle.InitCustomError(http.StatusBadRequest, err)
		wrappedErr := fmt.Errorf("UserRepoimple.Update: %w", cErr)		
		return user, wrappedErr
	}
	return user, nil
}

func (ur UserRepoimple) Delete(user *model.User) (*model.User, error) {
	db := ur.db

	err := db.Delete(&user).Error
	if err != nil {
		cErr := errorhandle.InitCustomError(http.StatusBadRequest, err)
		wrappedErr := fmt.Errorf("UserRepoimple.Delete: %w", cErr)		
		return nil, wrappedErr
	}
	return user, nil
}
