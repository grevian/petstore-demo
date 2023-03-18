package service

import (
	"context"
	"net/http"
	"os"

	openapi "github.com/grevian/petstore-demo/api/go"
)

var _ openapi.PetApiServicer = (*petAPIService)(nil)

type petAPIService struct {
	storage *petstoreStorage
}

func NewPetAPIService(storage *petstoreStorage) *petAPIService {
	return &petAPIService{
		storage: storage,
	}
}

func (p petAPIService) AddPet(ctx context.Context, pet openapi.Pet) (openapi.ImplResponse, error) {
	petId := p.storage.StorePet(&pet)
	pet.Id = petId
	return openapi.Response(http.StatusOK, pet), nil
}

func (p petAPIService) DeletePet(ctx context.Context, petId int64, apiKey string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPIService) FindPetsByStatus(ctx context.Context, status string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPIService) FindPetsByTags(ctx context.Context, tags []string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPIService) GetPetById(ctx context.Context, petId int64) (openapi.ImplResponse, error) {
	pet, err := p.storage.LoadPet(petId)
	if err != nil {
		return openapi.Response(http.StatusNotFound, nil), err
	}

	return openapi.Response(http.StatusOK, pet), nil
}

func (p petAPIService) UpdatePet(ctx context.Context, pet openapi.Pet) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPIService) UpdatePetWithForm(ctx context.Context, petId int64, petName string, petStatus string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPIService) UploadFile(ctx context.Context, petId int64, additionalMetadata string, file *os.File) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
