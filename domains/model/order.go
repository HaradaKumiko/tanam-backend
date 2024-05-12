package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	OrderPending    TransactionStatus = "pending"
	OrderCancel     TransactionStatus = "cancel"
	OrderExpire     TransactionStatus = "expire"
	OrderFailure    TransactionStatus = "failure"
	OrderSettlement TransactionStatus = "settlement"
)

type Order struct {
	OrderID           uuid.UUID         `gorm:"primaryKey;type:varchar(255);"`
	AuthID            uuid.UUID         `gorm:"type:varchar(255);nullable"`
	PlantID           uuid.UUID         `gorm:"type:varchar(255);nullable"`
	Qty               float32           `gorm:"type:numeric"`
	BasePrice         float64           `gorm:"type:decimal(10,2)"`
	TotalPrice        float64           `gorm:"type:decimal(10,2)"`
	PaymentType       string            `gorm:"type:varchar(255);nullable"`
	Token             string            `gorm:"type:varchar(255);nullable"`
	RedirectURL       string            `gorm:"type:varchar(255);nullable"`
	TransactionStatus TransactionStatus `gorm:"type:enum('pending','cancel', 'expire', 'failure', 'settlement');default:'pending'"`
	CreatedAt         time.Time         `gorm:"autoCreateTime"`
	UpdatedAt         time.Time         `gorm:"autoUpdateTime"`
}
