package repository

import (
	"errors"
	"tanam-backend/database"
	"tanam-backend/domains/model"
	"tanam-backend/domains/web/order"
	"tanam-backend/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
}

func InitOrderRepository() OrderRepository {
	return OrderRepository{}
}

func (p *OrderRepository) CreateOrderRepository(order order.CreateOrderMidtransRequest, token, redirectURL, authId string) (entities.Order, error) {

	newOrder := model.Order{
		OrderID:     uuid.MustParse(order.OrderID),
		AuthID:      uuid.MustParse(authId),
		PlantID:     uuid.MustParse(order.PlantID),
		Qty:         float32(order.ItemDetails.Qty),
		BasePrice:   float64(order.ItemDetails.Price),
		TotalPrice:  float64(order.Price),
		Token:       token,
		RedirectURL: redirectURL,
	}

	if err := database.DB.Create(&newOrder).Error; err != nil {
		return entities.Order{}, err
	}

	createdOrder := entities.Order{
		OrderID:     newOrder.OrderID,
		AuthID:      newOrder.AuthID,
		PlantID:     newOrder.PlantID,
		Qty:         int64(newOrder.Qty),
		BasePrice:   newOrder.BasePrice,
		TotalPrice:  newOrder.TotalPrice,
		Token:       newOrder.Token,
		RedirectURL: newOrder.RedirectURL,
		CreatedAt:   newOrder.CreatedAt,
		UpdatedAt:   newOrder.UpdatedAt,
	}

	return createdOrder, nil
}

func (p *OrderRepository) GetOrderByAuthIdRepository(authId string) ([]entities.Order, error) {
	var order []entities.Order
	if err := database.DB.Where("auth_id = ?", authId).Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, errors.New("order not found")
		}

		return order, err

	}
	return order, nil
}

func (p *OrderRepository) GetAllOrdersRepository() ([]entities.Order, error) {
	var orders []entities.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (p *OrderRepository) GetOrderByIdRepository(orderId string) (entities.Order, error) {
	var order entities.Order
	if err := database.DB.Where("order_id = ?", orderId).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, errors.New("order not found")
		}

		return order, err

	}
	return order, nil
}
