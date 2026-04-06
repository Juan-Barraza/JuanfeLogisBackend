package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string    `gorm:"size:100;not null"`
	Email        string    `gorm:"size:150;unique;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	Role         string    `gorm:"size:50;default:'admin'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
