package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Trijavico/ensolvers-assesment/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = db.AutoMigrate(&model.User{}, &model.Note{})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Completed migration!")
}
