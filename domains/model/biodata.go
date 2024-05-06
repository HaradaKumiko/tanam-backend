package model

import (
	"time"

	"github.com/google/uuid"
)

type Biodata struct {
	BiodataID   uuid.UUID `gorm:"type:varchar(100);"`
	AuthID      uuid.UUID `gorm:"type:varchar(100);nullable"`
	Fullname    string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
