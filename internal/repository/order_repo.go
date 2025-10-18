package repository

import (
	"database/sql"
	"fmt"
	"glassesbot/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save (order *domain.Order) error {

	res, err := r.db.Exec("INSERT INTO orders (frame, lenses, work_date, working_time) VALUES ($1, $2, $3, $4)",
		order.Frame, order.Lenses, order.Date, int64(order.WorkingTime.Minutes()))

	fmt.Println(res)

	return err
}