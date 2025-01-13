package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Trijavico/ensolvers-assesment/internal/model/dto"
	"github.com/Trijavico/ensolvers-assesment/internal/service"
)

type NoteController struct {
	noteService *service.NoteService
}

func NewNoteController(noteService *service.NoteService) *NoteController {
	return &NoteController{
		noteService: noteService,
	}
}

func (note *NoteController) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(dto.AuthUser)

	notes, err := note.noteService.GetAllUserNotes(user.ID)
	if err != nil {
		http.Error(w, "Unable to retrieve all notes", http.StatusInternalServerError)
		return
	}

	err = RespondJSON(w, r, http.StatusOK, notes)
	if err != nil {
		http.Error(w, "Unable to retrieve all notes", http.StatusInternalServerError)
	}
}

func (note *NoteController) GetAllNotesArchives(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(dto.AuthUser)

	notes, err := note.noteService.GetAllUserArchives(user.ID)
	if err != nil {
		http.Error(w, "Unable to retrieve all notes", http.StatusInternalServerError)
		return
	}

	err = RespondJSON(w, r, http.StatusOK, notes)
	if err != nil {
		http.Error(w, "Unable to retrieve all notes", http.StatusInternalServerError)
	}
}

func (note *NoteController) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")
	id, _ := strconv.Atoi(pathValue)

	value, err := note.noteService.GetNote(uint(id))
	if err != nil {
		http.Error(w, "Unable to retrieve note", http.StatusInternalServerError)
		return
	}

	err = RespondJSON(w, r, http.StatusOK, value)
	if err != nil {
		http.Error(w, "Unable to retrieve note", http.StatusInternalServerError)
	}
}

func (note *NoteController) CreateNote(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(dto.AuthUser)

	var createNote dto.CreateNote
	err := json.NewDecoder(r.Body).Decode(&createNote)
	if err != nil {
		http.Error(w, "Unable to create note", http.StatusInternalServerError)
		return
	}

	createNote.UserID = user.ID
	err = note.noteService.SaveNote(createNote)
	if err != nil {
		http.Error(w, "Unable to create note", http.StatusInternalServerError)
		return
	}

	err = RespondJSON(w, r, http.StatusCreated, "Created a note")
	if err != nil {
		http.Error(w, "Unable to create note", http.StatusInternalServerError)
	}
}

func (note *NoteController) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var updateNote dto.Note
	err := json.NewDecoder(r.Body).Decode(&updateNote)
	if err != nil {
		http.Error(w, "Unable to create note", http.StatusInternalServerError)
		return
	}

	err = note.noteService.UpdateNote(updateNote)
	if err != nil {
		http.Error(w, "Unable to create note", http.StatusInternalServerError)
		return
	}

	err = RespondJSON(w, r, http.StatusOK, "updated a note")
	if err != nil {
		http.Error(w, "Unable to update note", http.StatusInternalServerError)
	}
}

func (note *NoteController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")
	id, _ := strconv.Atoi(pathValue)

	err := note.noteService.DeleteNote(uint(id))
	if err != nil {
		http.Error(w, "Unable to delete note", http.StatusInternalServerError)
		return
	}

	err = RespondJSON(w, r, http.StatusOK, "deleted a note")
	if err != nil {
		http.Error(w, "Unable to delete note", http.StatusInternalServerError)
	}
}
