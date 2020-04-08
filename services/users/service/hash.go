package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
)

type Hasher interface {
	Hash(user *User) (string, error)
}

type UserHasher struct{}

func (h UserHasher) Hash(user *User) (string, error) {
	if len(user.Username) == 0 {
		return "", errors.New("cant get hash of empty username")
	}

	username := strings.TrimSpace(user.Username)
	username = strings.ToLower(username)

	hash := md5.Sum([]byte(username))

	return hex.EncodeToString(hash[:]), nil
}
