package repository

import (
	"database/sql"

	"github.com/savanyv/e-commerce-task-be/internal/models"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
