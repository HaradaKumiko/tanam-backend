package entities

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID           uuid.UUID `json:"order_id"`
	AuthID            uuid.UUID `json:"auth_id"`
	PlantID           uuid.UUID `json:"plant_id"`
	Qty               int64     `json:"qty"`
	BasePrice         float64   `json:"base_price"`
	TotalPrice        float64   `json:"total_price"`
	PaymentType       string    `json:"payment_type"`
	Token             string    `json:"token"`
	RedirectURL       string    `json:"redirect_url"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
