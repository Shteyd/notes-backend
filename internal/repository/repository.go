package repository

import (
	"github.com/Shteyd/notes-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type NotesRepository interface {
	CreateNote(userId int, note models.InputNote) (int, error)
	GetNoteByID(userId, noteId int) (models.Note, error)
	GetNotesByUserID(userId int) ([]models.Note, error)
	UpdateNote(userId, noteId int, note models.InputNote) error
	DeleteNote(userId, noteId int) error
}

type UsersRepository interface {
	CreateUser(user models.InputUser) (int, error)
	GetUserID(email, password string) (int, error)
}

type Repository struct {
	NotesRepository
	UsersRepository
}

func NewRepostitory(db *sqlx.DB) *Repository {
	return &Repository{
		NotesRepository: NewNotesRepo(db),
		UsersRepository: NewUsersRepo(db),
	}
}
