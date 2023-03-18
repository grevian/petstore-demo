package service

import (
	"context"
	"os"

	openapi "github.com/grevian/petstore-demo/api/go"
)

var _ openapi.PetApiServicer = (*petAPIService)(nil)

type petAPIService struct {
}

func NewPetAPIService() *petAPIService {
	return &petAPIService{}
}

func (p petAPIService) AddPet(ctx context.Context, pet openapi.Pet) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
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
