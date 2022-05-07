package service

import (
	"errors"
	"gorpc/pb"
	"sync"
)

var ErrAlreadyExist = errors.New("record already exists")
var ErrNotFound = errors.New("record doesn't exist")

type LaptopStore interface {
	Save(Laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

type DBLaptopStore struct {
}

func NewInMemLaptopStore() *InMemLaptopStore {
	return &InMemLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemLaptopStore) Save(Laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[Laptop.Id] != nil {
		return ErrAlreadyExist

	}
	store.data[Laptop.Id] = Laptop
	return nil
}

func (store *InMemLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	laptop := store.data[id]
	if laptop == nil {
		return nil, ErrNotFound
	}
	return laptop, nil
}
