package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthRole string

const (
	RoleDonor AuthRole = "donor"
	RoleAdmin AuthRole = "admin"
)

type Auth struct {
	AuthID    uuid.UUID `gorm:"type:varchar(100);"`
	Email     string    `gorm:"type:varchar(255);not null;unique"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Profile   string    `gorm:"type:varchar(255);default:'https://avatars.githubusercontent.com/u/42530587'"`
	Role      AuthRole  `gorm:"type:enum('donor','admin');default:'donor'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
