package storage

import "main/types"

type MemotyStorage struct {
}

func NewMemoryStorage() *MemotyStorage {
	return &MemotyStorage{}
}

func (s *MemotyStorage) Get(id string) *types.User {
	return &types.User{
		ID:   1,
		Name: "cuong",
	}
}
