package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Trijavico/ensolvers-assesment/internal/controller"
	"github.com/Trijavico/ensolvers-assesment/internal/database"
	"github.com/Trijavico/ensolvers-assesment/internal/middleware"
	"github.com/Trijavico/ensolvers-assesment/internal/repository"
	"github.com/Trijavico/ensolvers-assesment/internal/service"
	"github.com/joho/godotenv"
)

var (
	PORT int
)

func main() {
	flag.IntVar(&PORT, "port", 8080, "HTTP server port")

	flag.Parse()
	godotenv.Load()

	postgresDB, err := database.NewPostgresDB()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	authService := service.NewAuthService(
		repository.NewUserRepository(postgresDB),
		&service.JwtService{},
	)
	noteService := service.NewNoteService(
		repository.NewNoteRepository(postgresDB),
	)

	authController := controller.NewAuthController(authService)
	noteController := controller.NewNoteController(noteService)

	router := AddRoutes(
		authController,
		noteController,
	)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		Handler:      middleware.CORS(router),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	slog.Info(fmt.Sprintf("Server running on port: %d", PORT))
	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
