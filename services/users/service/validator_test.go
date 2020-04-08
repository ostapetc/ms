package service

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type registerValidatorMock struct {
	mock.Mock
}

func (r registerValidatorMock) Validate(users *Storage, user *User) error {
	args := r.Called(users, user)
	return args.Error(0)
}

func TestRegisterValidator_ValidateUserName(t *testing.T) {

}

func TestRegisterValidator_ValidatePassword(t *testing.T) {

}
