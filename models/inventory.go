package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	CreatedAt time.Time
}

type Box struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string    `gorm:"size:100;not null"`
	LocationID   uint
	GeneralLabel string `gorm:"size:150"`
	QRCodeURL    string `gorm:"type:text"`
	CreatedAt    time.Time

	Location Location `gorm:"foreignKey:LocationID"` // Relación GORM
}

type ProductType struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
}

type Donor struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"size:150;not null"`
	Type      string    `gorm:"size:50;not null"` // 'empresa' o 'persona'
	CreatedAt time.Time
}

type Product struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductTypeID     uint
	DonorID           uuid.UUID
	Size              string  `gorm:"size:20"`
	DonationPrice     float64 `gorm:"type:decimal(10,2);not null"`
	SalePrice         float64 `gorm:"type:decimal(10,2);not null"`
	PhysicalCondition string  `gorm:"size:50"`
	Disposition       string  `gorm:"size:50"`
	CreatedAt         time.Time

	ProductType ProductType `gorm:"foreignKey:ProductTypeID"`
	Donor       Donor       `gorm:"foreignKey:DonorID"`
}

type BoxStock struct {
	BoxID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Quantity  int       `gorm:"not null;default:0"`

	Box     Box     `gorm:"foreignKey:BoxID"`
	Product Product `gorm:"foreignKey:ProductID"`
}
