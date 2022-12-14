package services

import (
	"github.com/Shteyd/notes-backend/internal/models"
	"github.com/Shteyd/notes-backend/internal/repository"
)

type NoteService struct {
	repo repository.NotesRepository
}

func NewNoteService(repo repository.NotesRepository) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) CreateNote(userId int, note models.InputNote) (int, error) {
	return s.repo.CreateNote(userId, note)
}

func (s *NoteService) GetUserNotes(userId int) ([]models.Note, error) {
	return s.repo.GetNotesByUserID(userId)
}

func (s *NoteService) GetNoteByID(userId, noteId int) (models.Note, error) {
	return s.repo.GetNoteByID(userId, noteId)
}

func (s *NoteService) UpdateNote(userId, noteId int, note models.InputNote) error {
	return s.repo.UpdateNote(userId, noteId, note)
}

func (s *NoteService) DeleteNote(userId, noteId int) error {
	return s.repo.DeleteNote(userId, noteId)
}
