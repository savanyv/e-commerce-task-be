package repository

import (
	"database/sql"

	"github.com/savanyv/e-commerce-task-be/internal/models"
)

type OrderRepository interface {
	Create(order *models.Order) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(order *models.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var orderID int
	err = tx.QueryRow(`INSERT INTO orders (user_id) VALUES ($1) RETURNING id`, order.UserID).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range order.Items {
		_, err := tx.Exec(`INSERT INTO order_items (order_id, variant_id, quantity, price) VALUES ($1, $2, $3, $4)`, orderID, item.VariantID, item.Quantity, item.Price)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(`UPDATE product_variants SET stock = stock - $1 WHERE id = $2 AND stock >= $1`, item.Quantity, item.VariantID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	order.ID = orderID
	return nil
}

