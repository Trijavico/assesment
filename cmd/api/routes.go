package main

import (
	"fmt"
	"net/http"

	"github.com/Trijavico/ensolvers-assesment/internal/controller"
	"github.com/Trijavico/ensolvers-assesment/internal/middleware"
)

func AddRoutes(
	authController *controller.AuthController,
	noteController *controller.NoteController,
) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", handleIndex)

	router.HandleFunc("POST /api/auth/signup", authController.CreateAccount)
	router.HandleFunc("POST /api/auth/login", authController.LogInAccount)

	notesRoutes := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", middleware.Authenticated(notesRoutes)))

	notesRoutes.Handle("GET /notes", makeHandler(noteController.GetAllNotes))
	notesRoutes.Handle("POST /notes", makeHandler(noteController.CreateNote))
	notesRoutes.Handle("GET /notes/archives", makeHandler(noteController.GetAllNotesArchives))
	notesRoutes.Handle("GET /notes/{id}", makeHandler(noteController.GetNoteByID))
	notesRoutes.Handle("PATCH /notes", makeHandler(noteController.UpdateNote))
	notesRoutes.Handle("DELETE /notes/{id}", makeHandler(noteController.DeleteNote))

	return router
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		message := fmt.Sprintf("NOT FOUND url: %s", r.URL)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	w.Write([]byte("hello from http server"))
}
