package usecase

import (
	"go-trial/domain/model"
	"go-trial/domain/repo"
)

type UserUsecase struct {
	userRepo repo.UserRepo
}

func SetUserUsecase(ur repo.UserRepo) UserUsecase {
	// ur には userRepoimple が注入される
	return UserUsecase{userRepo: ur}
}

func (uu UserUsecase) GetList() ([]model.User, error) {
	// uu.userRepo には userRepoimple が注入されているため
	// userRepoimple.GetList() が実行される
	users, err := uu.userRepo.GetList()
	if err != nil {
		return nil, err
	}
	return users, err
}