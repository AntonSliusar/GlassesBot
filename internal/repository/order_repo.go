package repository

import (
	"database/sql"
	"glassesbot/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepositiry(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save (order *domain.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (frame, lensses, date, working_time) VALUES ($1, $2, $3, $4)",
		order.Frame, order.Lenses, order.Date, order.WorkingTime)

	return err
}