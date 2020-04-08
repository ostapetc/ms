package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsersService_Register(t *testing.T) {
	storage := NewStorage()
	hasher := UserHasher{}
	user := NewUser(1, "user", "password", hasher)

	validator := registerValidatorMock{}
	validator.On("Validate", storage, user).Return(nil)

	service := NewUsersService(storage, validator)
	authKey, err := service.Register(user)

	assertUserRegisterSuccess(t, service, user, authKey, err)
}

func TestUsersService_RegisterFull(t *testing.T) {
	storage := NewStorage()
	validator := RegisterValidator{}
	hasher := UserHasher{}

	service := NewUsersService(storage, validator)
	user := NewUser(1, "user", "password", hasher)

	authKey, err := service.Register(user)

	assertUserRegisterSuccess(t, service, user, authKey, err)
}

func assertUserRegisterSuccess(t *testing.T, service *UsersService, user *User, authKey string, err error) {
	assert.NotEmpty(t, authKey)
	assert.Nil(t, err)

	dbUser := service.GetByID(1)

	assert.Equal(t, dbUser.Username, user.Username)
	assert.Equal(t, dbUser.Password, user.Password)
	assert.Equal(t, dbUser.Hash, user.Hash)
	assert.NotEmpty(t, dbUser.Hash)
}
