package repository

import (
	"github.com/Shteyd/notes-backend/internal/database"
	"github.com/Shteyd/notes-backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) CreateUser(user models.InputUser) (int, error) {
	var id int

	row := r.db.QueryRow(database.CreateNewUser, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *UsersRepo) GetUserID(email, password string) (int, error) {
	var id int

	row := r.db.QueryRow(database.GetUserById, email, password)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}
