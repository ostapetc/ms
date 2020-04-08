package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type userHasherMock struct {
	mock.Mock
}

func (r registerValidatorMock) Hash(users *Storage, user *User) error {
	args := r.Called(users, user)
	return args.Error(0)
}

type testCase struct {
	Username1 string
	Username2 string
	EqualHash bool
}

func TestUserHashCreator_CreateHash(t *testing.T) {
	tcases := []*testCase{
		&testCase{
			Username1: "name",
			Username2: "name",
			EqualHash: true,
		},
		&testCase{
			Username1: "name",
			Username2: "Name",
			EqualHash: true,
		},
		&testCase{
			Username1: "name ",
			Username2: "Name",
			EqualHash: true,
		},
		//&testCase{
		//	Username1: "first",
		//	Username2: "second",
		//	EqualHash: false,
		//},
	}

	hasher := UserHasher{}

	for _, tcase := range tcases {
		user1 := &User{Username: tcase.Username1}
		hash1, err := hasher.Hash(user1)
		assert.Nil(t, err)
		assert.NotEmpty(t, hash1)

		user2 := &User{Username: tcase.Username2}
		hash2, err := hasher.Hash(user2)
		assert.Nil(t, err)
		assert.NotEmpty(t, hash2)

		if tcase.EqualHash {
			assert.Equal(t, hash1, hash2)
		} else {
			assert.NotEqual(t, hash1, hash2)
		}
	}
}

func TestUser_GetHashError(t *testing.T) {
	hasher := UserHasher{}
	_, err := hasher.Hash(&User{})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "cant get hash of empty username")
}
