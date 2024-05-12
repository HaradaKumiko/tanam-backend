package admin

import (
	"tanam-backend/entities"
	"tanam-backend/repository"
)

type OrderService struct {
	repository repository.OrderRepository
}

func InitOrderService() OrderService {
	return OrderService{
		repository: repository.InitOrderRepository(),
	}
}

func (service *OrderService) GetAllOrderService() ([]entities.Order, error) {
	orders, err := service.repository.GetAllOrdersRepository()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (service *OrderService) GetOrderByIdService(orderId string) (entities.Order, error) {
	order, err := service.repository.GetOrderByIdRepository(orderId)
	if err != nil {
		return order, err
	}

	return order, nil
}
