package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	db, err := gorm.Open(postgres.Open(connStr), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, err
}
