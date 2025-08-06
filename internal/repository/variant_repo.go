package repository

import (
	"database/sql"

	"github.com/savanyv/e-commerce-task-be/internal/models"
)

type VariantRepository interface {
	CreateBulk(productID int, variants []*models.Variant) error
	FindByProductID(productID int) ([]*models.Variant, error)
}

type variantRepository struct {
	db *sql.DB
}

func NewVariantRepository(db *sql.DB) VariantRepository {
	return &variantRepository{
		db: db,
	}
}

func (r *variantRepository) CreateBulk(productID int, variatns []*models.Variant) error {
	query := `INSERT INTO product_variants (product_id, name, price, stock) VALUES ($1, $2, $3, $4)`
	for _, v := range variatns {
		_, err := r.db.Exec(query, productID, v.Name, v.Price, v.Stock)
		if err != nil {
			return err
		}
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
