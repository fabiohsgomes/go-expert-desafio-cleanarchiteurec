package graph

import "github.com/fabiohsgomes/go-expert-desafio-cleanarchiteurec/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrderUseCase   *usecase.ListOrderUseCase
}
