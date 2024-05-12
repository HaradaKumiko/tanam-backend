package entities

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	AuthID    uuid.UUID `json:"auth_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Profile   string    `json:"profile"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
