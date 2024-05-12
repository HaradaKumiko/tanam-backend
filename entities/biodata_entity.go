package entities

import (
	"time"

	"github.com/google/uuid"
)

type Biodata struct {
	AuthID      uuid.UUID `json:"auth_id"`
	Fullname    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
