package model

import "errors"

type User struct {
	Base
	Name string
}

func InitUser(name string) (*User, error) {
	if name == "" {
		err := errors.New("name is required")
		return nil, err
	}
	return &User{Name: name}, nil
}

func (u *User) ChangeField(name string) (error) {
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