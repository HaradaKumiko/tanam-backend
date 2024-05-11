package entities

import (
	"time"

	"github.com/google/uuid"
)

type Plant struct {
	PlantID     uuid.UUID `json:"plant_id"`
	Name        string    `json:"name" `
	Description string    `json:"description"`
	Picture     string    `json:"picture"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
