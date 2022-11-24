package model

import "errors"

type User struct {
	Base
	Name string
}

func InitUser(name string) (User, error) {
	if name == "" {
		err := errors.New("name is required")
		return User{}, err
	}
	user := User{Name: name}
	return user, nil
}

// 構造体の中身を操作時: 引数の構造体を &(ポインタ)で投げて *(変数)で受け取る
func ChangeField(u *User, name string) (error) {
	var err error
	if (u == nil) {
        err = errors.New("user is required")
		return err
    }
	if name == "" {
		err = errors.New("name is required")
		return err
	}
	u.Name = name
	return nil
}