package storage

import (
	"fmt"
	"godis/internal/storage/cache"
	"sync"
)

type Storage struct {
	KeyValueStorage *sync.Map
	StoredCounter   int
}

func keyFormatToString(key any) string {
	return fmt.Sprint("", key)

}

func NewStorage(initStorage *sync.Map) Storage {
	storage := Storage{KeyValueStorage: initStorage, StoredCounter: 0}
	return storage
}

func (s *Storage) Set(key, val any) error {
	s.KeyValueStorage.Store(key, val)
	err := cache.WriteDataToCache(keyFormatToString(key), val)
	if err != nil {
		return err
	}
	s.StoredCounter += 1
	return nil
}

func (s *Storage) Get(key any) any {
	data, ok := s.KeyValueStorage.Load(key)
	if ok == false {
		data = cache.GetDataFromCache(keyFormatToString(key))
	}
	return data
}
