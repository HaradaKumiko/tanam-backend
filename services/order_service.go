package services

import (
	"tanam-backend/domains/web/order"
	"tanam-backend/entities"
	"tanam-backend/repository"
	"tanam-backend/utils/midtrans"

	"github.com/google/uuid"
)

type OrderService struct {
	plantRepository   repository.PlantRepository
	orderRepository   repository.OrderRepository
	authRepository    repository.AuthRepository
	biodataRepository repository.BiodataRepository
}

func InitOrderService() OrderService {
	return OrderService{
		plantRepository:   repository.InitPlantRepository(),
		orderRepository:   repository.InitOrderRepository(),
		authRepository:    repository.InitAuthRepository(),
		biodataRepository: repository.InitBiodataRepository(),
	}
}

func (service *OrderService) CreateOrderService(request order.CreateOrderRequest, authId string) (any, error) {
	plant, err := service.plantRepository.GetPlantByIdRepository(request.PlantID)
	if err != nil {
		return plant, err
	}

	auth, err := service.authRepository.FindAuthById(authId)
	if err != nil {
		return auth, err
	}

	biodata, err := service.biodataRepository.FIndBiodataByAuthId(authId)
	if err != nil {
		return biodata, err
	}

	amount := plant.Price * float64(request.Quantity)

	orderData := order.CreateOrderMidtransRequest{
		PlantID:        plant.PlantID.String(),
		OrderID:        uuid.NewString(),
		Price:          int64(amount),
		CustomerDetail: order.CustomerDetail{FName: biodata.Fullname, Email: auth.Email, Phone: biodata.PhoneNumber},
		ItemDetails:    order.ItemDetails{ID: plant.PlantID.String(), Name: plant.Name, Price: int64(plant.Price), Qty: request.Quantity},
	}

	midtransToken, _ := midtrans.NewMidtransPayment().CreateTransaction(orderData)

	dataOrder, err := service.orderRepository.CreateOrderRepository(orderData, midtransToken.Token, midtransToken.RedirectURL, authId)
	if err != nil {
		return dataOrder, err
	}

	return dataOrder, nil
}

func (service *OrderService) GetOrderDonorService(authId string) ([]entities.Order, error) {
	order, err := service.orderRepository.GetOrderByAuthIdRepository(authId)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (service *OrderService) GetOrderByIdService(orderId string) (entities.Order, error) {
	order, err := service.orderRepository.GetOrderByIdRepository(orderId)
	if err != nil {
		return order, err
	}

	return order, nil
}

