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

func (uu UserUsecase) GetSingle(id string) (model.User, error) {
	// uu.userRepo には userRepoimple が注入されているため
	// userRepoimple.GetSingle() が実行される
	user, err := uu.userRepo.GetSingle(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uu UserUsecase) Create(name string) (model.User, error) {
	user, err := model.SetNewUser(name)
	if err != nil {
		return user, err
	}
	// uu.userRepo には userRepoimple が注入されているため
	// userRepoimple.Create() が実行される
	createdUser, err := uu.userRepo.Create(user)
	if err != nil {
		return createdUser, err
	}
	return createdUser, nil
}

func (uu UserUsecase) Update(id string, name string) (model.User, error) {
	user, err := uu.userRepo.GetSingle(id)
	if err != nil {
		return user, err
	}
	user, err = user.SetUser(name)
	if err != nil {
		return user, err
	}
	// uu.userRepo には userRepoimple が注入されているため
	// userRepoimple.Update() が実行される
	updatedUser, err := uu.userRepo.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (uu UserUsecase) Delete(id string) (model.User, error) {
	user, err := uu.userRepo.GetSingle(id)
	if err != nil {
		return user, err
	}
	// uu.userRepo には userRepoimple が注入されているため
	// userRepoimple.Delete() が実行される
	deletedUser, err := uu.userRepo.Delete(user)
	if err != nil {
		return deletedUser, err
	}
	return deletedUser, nil
}
