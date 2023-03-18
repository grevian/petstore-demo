package service

import (
	"context"

	openapi "github.com/grevian/petstore-demo/api/go"
)

var _ openapi.StoreApiServicer = (*storeAPIService)(nil)

type storeAPIService struct {
}

func NewStoreApiService() *storeAPIService {
	return &storeAPIService{}
}

func (s storeAPIService) DeleteOrder(ctx context.Context, orderId int64) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s storeAPIService) GetInventory(ctx context.Context) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s storeAPIService) GetOrderById(ctx context.Context, orderId int64) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s storeAPIService) PlaceOrder(ctx context.Context, order openapi.Order) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
