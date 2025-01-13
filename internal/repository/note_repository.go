package repository

import (
	"context"

	"github.com/Trijavico/ensolvers-assesment/internal/model"
	"github.com/Trijavico/ensolvers-assesment/internal/model/dto"
	"gorm.io/gorm"
)

type NoteRepository interface {
	FindAllByUser(ctx context.Context, userID uint) ([]model.Note, error)
	FindByID(context.Context, uint) (model.Note, error)
	FindAllArchived(context.Context, uint) ([]model.Note, error)

	Save(context.Context, *model.Note) error
	Update(context.Context, dto.Note) error
	Delete(context.Context, uint) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{
		db: db,
	}
}

func (noteRepo *noteRepository) FindAllByUser(ctx context.Context, userID uint) ([]model.Note, error) {
	var notes []model.Note

	err := noteRepo.db.WithContext(ctx).Where("user_id = ? AND archived = ?", userID, false).Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (noteRepo *noteRepository) FindAllArchived(ctx context.Context, userID uint) ([]model.Note, error) {
	var notes []model.Note

	err := noteRepo.db.WithContext(ctx).Where("user_id = ? AND archived = ?", userID, true).Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (noteRepo *noteRepository) FindByID(ctx context.Context, id uint) (model.Note, error) {
	var note model.Note

	err := noteRepo.db.WithContext(ctx).First(&note, id).Error
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}

func (noteRepo *noteRepository) Save(ctx context.Context, note *model.Note) error {
	err := noteRepo.db.WithContext(ctx).Create(note).Error
	if err != nil {
		return err
	}

	return nil
}

func (noteRepo *noteRepository) Update(ctx context.Context, note dto.Note) error {
	err := noteRepo.db.WithContext(ctx).Model(&model.Note{}).First(&model.Note{}, note.ID).Error
	if err != nil {
		return err
	}

	err = noteRepo.db.WithContext(ctx).Model(&model.Note{ID: note.ID}).Save(&note).Error
	if err != nil {
		return err
	}

	return nil
}

func (noteRepo *noteRepository) Delete(ctx context.Context, id uint) error {
	err := noteRepo.db.WithContext(ctx).Delete(&model.Note{ID: id}).Error
	if err != nil {
		return err
	}

	return nil
}
