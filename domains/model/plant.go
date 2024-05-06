package model

import (
	"time"

	"github.com/google/uuid"
)

type Plant struct {
	PlantID     uuid.UUID `gorm:"type:varchar(100);"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(20);not null"`
	Picture     string    `gorm:"type:varchar(20);not null"`
	Price       float64   `gorm:"type:decimal(10,2)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
