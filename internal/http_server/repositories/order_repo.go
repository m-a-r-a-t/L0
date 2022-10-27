package repositories

import "database/sql"

type orderRepo struct {
	db *sql.DB
}

func (o *orderRepo) GetOrderById(string) {
	// o.db
}

func NewOrderRepo(db *sql.DB) *orderRepo {
	return &orderRepo{db: db}
}
