package config

import (
	"fmt"
	"juanfeLogis/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error conectando a la base de datos: \n", err)
	}

	log.Println("Conexión exitosa a PostgreSQL")

	DB = db
}

func MigrateDB() {
	if DB == nil {
		log.Fatal("Erro to execute migration")
	}
	db := DB
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	log.Println("Ejecutando AutoMigrate...")
	err := db.AutoMigrate(
		&models.User{},
		&models.Location{},
		&models.Box{},
		&models.ProductType{},
		&models.Donor{},
		&models.Product{},
		&models.BoxStock{},
		&models.Transaction{},
		&models.TransactionItem{},
	)

	if err != nil {
		log.Fatal("Error durante la migración: \n", err)
	}
	log.Println("Tablas sincronizadas correctamente")
}
