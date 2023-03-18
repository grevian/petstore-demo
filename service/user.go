package service

import (
	"context"

	openapi "github.com/grevian/petstore-demo/api/go"
)

var _ openapi.UserApiServicer = (*userAPIService)(nil)

type userAPIService struct {
}

func NewUserApiService() *userAPIService {
	return &userAPIService{}
}

func (u userAPIService) CreateUser(ctx context.Context, user openapi.User) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userAPIService) CreateUsersWithListInput(ctx context.Context, users []openapi.User) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userAPIService) DeleteUser(ctx context.Context, username string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userAPIService) GetUserByName(ctx context.Context, username string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userAPIService) LoginUser(ctx context.Context, username string, password string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userAPIService) LogoutUser(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userAPIService) UpdateUser(ctx context.Context, username string, user openapi.User) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
