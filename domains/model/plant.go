package model

import (
	"time"

	"github.com/google/uuid"
)

type StatusPlant string

const (
	StatusAvailable   StatusPlant = "available"
	StatusUnavailable StatusPlant = "unavailable"
)

type Plant struct {
	PlantID     uuid.UUID `gorm:"primaryKey;type:varchar(255);"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Picture     string    `gorm:"type:varchar(255);not null"`
	Price       float64   `gorm:"type:decimal(10,2)"`
	Status      string    `gorm:"type:enum('available','unavailable');default:'available'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
