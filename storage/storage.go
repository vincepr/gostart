package storage

import "apiServer/types"

type Storage interface {
	Get(int) *types.User
}