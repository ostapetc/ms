package service

import (
	"errors"
)

const usernameMinLen = 3
const passwordMinLen = 6

type Validator interface {
	Validate(storage *Storage, user *User) error
}

type RegisterValidator struct{}

func (v RegisterValidator) Validate(users *Storage, user *User) error {
	if len(user.Username) < usernameMinLen {
		return errors.New("invalid user name length")
	}

	if len(user.Password) < passwordMinLen {
		return errors.New("invalid password length")
	}

	if _, ok := users.hashmap[user.Hash]; ok {
		return errors.New("username already exists")
	}

	if _, ok := users.idmap[user.ID]; ok {
		return errors.New("id already exists")
	}

	return nil
}
