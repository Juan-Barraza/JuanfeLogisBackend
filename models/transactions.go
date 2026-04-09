package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Type      string    `gorm:"size:50;not null"` // 'entrada', 'salida', 'devolucion', 'ajuste'
	UserID    uuid.UUID
	CreatedAt time.Time

	User  User              `gorm:"foreignKey:UserID"`
	Items []TransactionItem `gorm:"foreignKey:TransactionID"`
}

type TransactionItem struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TransactionID uuid.UUID
	ProductID     uuid.UUID
	BoxID         uuid.UUID
	Quantity      int       `gorm:"not null"`
	AppliedPrice  float64   `gorm:"type:decimal(10,2);not null"`

	Transaction Transaction `gorm:"foreignKey:TransactionID"`
	Product     Product     `gorm:"foreignKey:ProductID"`
	Box         Box         `gorm:"foreignKey:BoxID"`
}