package repository

import (
	"github.com/Shteyd/notes-backend/internal/database"
	"github.com/Shteyd/notes-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type NotesRepo struct {
	db *sqlx.DB
}

func NewNotesRepo(db *sqlx.DB) *NotesRepo {
	return &NotesRepo{db: db}
}

func (r *NotesRepo) CreateNote(userId int, note models.InputNote) (int, error) {
	query := database.CreateNoteWithoutContent
	args := append(make([]interface{}, 0), userId, note.Title)

	if note.Content != nil {
		query = database.CreateNote
		args = append(args, *note.Content)
	}

	var id int

	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *NotesRepo) GetNoteByID(userId, noteId int) (models.Note, error) {
	var res models.Note

	if err := r.db.Get(&res, database.GetNoteByID, noteId, userId); err != nil {
		return res, err
	}

	return res, nil
}

func (r *NotesRepo) GetNotesByUserID(userId int) ([]models.Note, error) {
	var res []models.Note

	if err := r.db.Get(&res, database.GetAllNotesByUserID, userId); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *NotesRepo) UpdateNote(userId, noteId int, note models.InputNote) error {
	var query string
	args := make([]interface{}, 0)

	if note.Title != nil && note.Content != nil {
		query = database.UpdateNote
		args = append(args, *note.Title, *note.Content, noteId, userId)
	} else if note.Title != nil {
		query = database.UpdateNoteWithoutContent
		args = append(args, *note.Title, noteId, userId)
	} else if note.Content != nil {
		query = database.UpdateNoteWithoutTitle
		args = append(args, *note.Content, noteId, userId)
	}

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *NotesRepo) DeleteNote(userId, noteId int) error {
	_, err := r.db.Exec(database.DeleteNoteByID, noteId, userId)
	return err
}
