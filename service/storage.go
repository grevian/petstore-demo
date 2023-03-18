package service

import (
	"errors"

	openapi "github.com/grevian/petstore-demo/api/go"
)

var ErrPetNotFound = errors.New("petstoreStorage: Not Found")

type petstoreStorage struct {
	indexCounter int64
	inMemory     map[int64]*openapi.Pet
}

func NewPetstoreStorage() *petstoreStorage {
	return &petstoreStorage{
		indexCounter: 0,
		inMemory:     make(map[int64]*openapi.Pet),
	}
}

func (s *petstoreStorage) StorePet(pet *openapi.Pet) int64 {
	petId := s.indexCounter
	s.indexCounter += 1
	s.inMemory[petId] = pet
	return petId
}

func (s *petstoreStorage) LoadPet(id int64) (*openapi.Pet, error) {
	if pet, exists := s.inMemory[id]; exists == true {
		return pet, nil
	}
	return nil, ErrPetNotFound
}
