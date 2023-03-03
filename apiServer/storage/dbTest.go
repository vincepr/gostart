package storage

import "apiServer/types"



type MemoryStorage struct{}

func NewMemordyStorage() *MemoryStorage{
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.User{
	return &types.User{
		ID: 1,
		Name: "Paul",
	}
}