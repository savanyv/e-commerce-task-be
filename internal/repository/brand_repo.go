package repository

import (
	"database/sql"

	"github.com/savanyv/e-commerce-task-be/internal/models"
)

type BrandRepository interface {
	FindAll() ([]*models.Brand, error)
	Create(brand *models.Brand) error
}

type brandRepository struct {
	db *sql.DB
}

func NewBrandRepository(db *sql.DB) BrandRepository {
	return &brandRepository{
		db: db,
	}
}

func (r *brandRepository) FindAll() ([]*models.Brand, error) {
	query := `SELECT id, name FROM brands`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []*models.Brand
	for rows.Next() {
		var b models.Brand
		if err := rows.Scan(&b.ID, &b.Name); err != nil {
			return nil, err
		}
		brands = append(brands, &b)
	}

	return brands, nil
}

func (r *brandRepository) Create(brand *models.Brand) error {
	query := `INSERT INTO brands (name) VALUES ($1) RETURNING id`
	if err := r.db.QueryRow(query, brand.Name).Scan(&brand.ID); err != nil {
		return err
	}

	return nil
}

