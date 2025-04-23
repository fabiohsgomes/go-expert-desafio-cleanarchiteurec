package usecase

import "github.com/fabiohsgomes/go-expert-desafio-cleanarchiteurec/internal/entity"


type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ListOrderUseCase) Execute() ([]*OrderOutputDTO, error) {
	orders, err := u.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var orderOutputDTOs []*OrderOutputDTO
	for _, order := range orders {
		orderOutputDTOs = append(orderOutputDTOs, &OrderOutputDTO{
			ID:     order.ID,
			Price:  order.Price,
			Tax:    order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orderOutputDTOs, nil
}
