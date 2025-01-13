package service

import (
	"context"

	"github.com/Trijavico/ensolvers-assesment/internal/model"
	"github.com/Trijavico/ensolvers-assesment/internal/model/dto"
	"github.com/Trijavico/ensolvers-assesment/internal/repository"
)

type NoteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) *NoteService {
	return &NoteService{
		noteRepo: noteRepo,
	}
}

func (service *NoteService) GetAllUserNotes(userID uint) ([]dto.Note, error) {
	result, err := service.noteRepo.FindAllByUser(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return []dto.Note{}, nil
	}

	notes := make([]dto.Note, len(result))

	for i, note := range result {
		notes[i] = dto.Note{
			ID:       note.ID,
			Title:    note.Title,
			Content:  note.Content,
			Archived: note.Archived,
			UserID:   note.UserID,
		}
	}

	return notes, nil
}

func (service *NoteService) GetAllUserArchives(userID uint) ([]dto.Note, error) {
	result, err := service.noteRepo.FindAllArchived(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return []dto.Note{}, nil
	}

	notes := make([]dto.Note, len(result))

	for i, note := range result {
		notes[i] = dto.Note{
			ID:       note.ID,
			Title:    note.Title,
			Content:  note.Content,
			Archived: note.Archived,
			UserID:   note.UserID,
		}
	}

	return notes, nil
}

func (service *NoteService) GetNote(id uint) (dto.Note, error) {
	result, err := service.noteRepo.FindByID(context.TODO(), id)
	if err != nil {
		return dto.Note{}, nil
	}

	var note dto.Note
	note.ID = result.ID
	note.Title = result.Title
	note.Content = result.Content
	note.Archived = result.Archived
	note.UserID = result.UserID

	return note, nil
}

func (service *NoteService) SaveNote(createNote dto.CreateNote) error {
	note := &model.Note{
		Title:    createNote.Title,
		Content:  createNote.Content,
		Archived: createNote.Archived,
		UserID:   createNote.UserID,
	}

	err := service.noteRepo.Save(context.Background(), note)
	if err != nil {
		return err
	}

	return nil
}

func (service *NoteService) UpdateNote(updateNote dto.Note) error {
	err := service.noteRepo.Update(context.Background(), updateNote)
	if err != nil {
		return err
	}

	return nil
}

func (service *NoteService) DeleteNote(id uint) error {
	err := service.noteRepo.Delete(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}
