package service

import (
	"sync"
)

type Storage struct {
	mx      sync.Mutex
	hashmap map[string]*User
	idmap   map[int64]*User
}

func NewStorage() *Storage {
	users := &Storage{}
	users.hashmap = make(map[string]*User)
	users.idmap = make(map[int64]*User)

	return users
}

type UsersService struct {
	storage           *Storage
	registerValidator Validator
	hashCreator       *Hasher
}

func NewUsersService(users *Storage, registerValidator Validator) *UsersService {
	return &UsersService{
		storage:           users,
		registerValidator: registerValidator,
	}
}

func (s *UsersService) Register(user *User) (authKey string, err error) {
	s.storage.mx.Lock()
	defer s.storage.mx.Unlock()

	err = s.registerValidator.Validate(s.storage, user)
	if err != nil {
		return "", err
	}

	s.storage.hashmap[user.Hash] = user
	s.storage.idmap[user.ID] = user

	return user.Hash, nil
}

func (s UsersService) GetByID(id int64) *User {
	user, _ := s.storage.idmap[id]
	return user
}
