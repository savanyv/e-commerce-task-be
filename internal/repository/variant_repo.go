package repository

import (
	"database/sql"

	"github.com/savanyv/e-commerce-task-be/internal/models"
)

type VariantRepository interface {
	CreateBulk(productID int, variants []*models.Variant) error
	FindByProductID(productID int) ([]*models.Variant, error)
	FindByID(ID int) (*models.Variant, error)
}

type variantRepository struct {
	db *sql.DB
}

func NewVariantRepository(db *sql.DB) VariantRepository {
	return &variantRepository{
		db: db,
	}
}

func (r *variantRepository) CreateBulk(productID int, variants []*models.Variant) error {
	for _, v := range variants {
		query := `INSERT INTO product_variants (product_id, name, price, strock) VALUES ($1, $2, $3, $4) RETURNING id`
		if err := r.db.QueryRow(query, productID, v.Name, v.Price, v.Stock).Scan(&v.ID); err != nil {
			return err
		}
		v.ProductID = productID
	}
	
	return nil
}

func (r *variantRepository) FindByProductID(productID int) ([]*models.Variant, error) {
	query := `SELECT id, product_id, name, price, stock FROM product_variants WHERE product_id = $1`
	rows, err := r.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []*models.Variant
	for rows.Next() {
		var v models.Variant
		if err := rows.Scan(&v.ID, &v.ProductID, &v.Name, &v.Price, &v.Stock); err != nil {
			return nil, err
		}
		variants = append(variants, &v)
	}

	return variants, nil
}

func (r *variantRepository) FindByID(ID int) (*models.Variant, error) {
	query := `SELECT id, product_id, name, price, stock FROM product_variants WHERE id = $1`
	var v models.Variant
	if err := r.db.QueryRow(query, ID).Scan(&v.ID, &v.ProductID, &v.Name, &v.Price, &v.Stock); err != nil {
		return nil, err
	}
	return &v, nil
}
