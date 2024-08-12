package storage

import "main/types"

type Storage interface {
	Get(string) *types.User
}
