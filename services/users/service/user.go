package service

import "strings"

type User struct {
	ID       int64
	Username string
	Password string
	Hash     string
}

func NewUser(id int64, username string, password string, hasher Hasher) *User {
	user := &User{
		ID:       id,
		Username: strings.TrimSpace(username),
		Password: password,
	}

	hash, err := hasher.Hash(user)
	if err != nil {
		panic(err)
	}

	user.Hash = hash

	return user
}
