package model

import "errors"

type User struct {
	Base
	Name string
}

func SetNewUser(name string) (User, error) {
	if name == "" {
		err := errors.New("name is required")
		return User{}, err
	}
	user := User{Name: name}
	return user, nil
}