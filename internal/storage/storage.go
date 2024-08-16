package storage

import (
	"fmt"
	"sync"
)

type Storage struct {
	KeyValueStorage *sync.Map
	StoredCounter   int
}

func NewStorage(initStorage *sync.Map) Storage {
	storage := Storage{KeyValueStorage: initStorage, StoredCounter: 0}
	return storage
}

func (s *Storage) Set(key, val any) {
	s.KeyValueStorage.Store(key, val)
	s.StoredCounter += 1
}

func (s *Storage) Get(key any) any {
	data, ok := s.KeyValueStorage.Load(key)
	if ok == false {
		return nil
	}
	fmt.Println(data)
	return data
}
