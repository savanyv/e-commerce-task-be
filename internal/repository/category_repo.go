package repository

import (
	"database/sql"

	"github.com/savanyv/e-commerce-task-be/internal/models"
)

type CategoryRepository interface {
	FindAll() ([]*models.Category, error)
	Create(category *models.Category) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) FindAll() ([]*models.Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}

	return categories, nil
}

func (r *categoryRepository) Create(category *models.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	if err := r.db.QueryRow(query, category.Name).Scan(&category.ID); err != nil {
		return err
	}

	return nil
}
