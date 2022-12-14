package services

import (
	"github.com/Shteyd/notes-backend/internal/models"
	"github.com/Shteyd/notes-backend/internal/repository"
	"github.com/Shteyd/notes-backend/pkg/env"
)

type Authorization interface {
	CreateUser(user models.InputUser) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Note interface {
	CreateNote(userId int, note models.InputNote) (int, error)
	GetUserNotes(userId int) ([]models.Note, error)
	GetNoteByID(userId, noteId int) (models.Note, error)
	UpdateNote(userId, noteId int, note models.InputNote) error
	DeleteNote(userId, noteId int) error
}

type Service struct {
	Authorization
	Note
}

func NewService(config *env.EnvConfig, repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(config, repos.UsersRepository),
		Note:          NewNoteService(repos.NotesRepository),
	}
}
