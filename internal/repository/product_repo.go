package repository

import (
	"database/sql"

	"github.com/savanyv/e-commerce-task-be/internal/models"
)

type ProductRepository interface {
	FindAll() ([]*models.Product, error)
	FindByID(id int) (*models.Product, error)
	Create(product *models.Product) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindAll() ([]*models.Product, error) {
	query := `SELECT id, name, type, category_id, brand_id FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.CategoryID, &p.BrandID); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}

func (r *productRepository) FindByID(id int) (*models.Product, error) {
	query := `SELECT id, name, type, category_id, brand_id FROM products WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var p models.Product
	if err := row.Scan(&p.ID, &p.Name, &p.Type, &p.CategoryID, &p.BrandID); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepository) Create(product *models.Product) error {
	query := `INSERT INTO products (name, type, catebrand_igory_id, d) VALUES ($1, $2, $3, $4) RETURNING id`
	if err := r.db.QueryRow(query, product.Name, product.Type, product.CategoryID, product.BrandID).Scan(&product.ID); err != nil {
		return err
	}

	return nil
}
