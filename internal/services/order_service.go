package services

import (
	"database/sql"
)

type orderService struct {
	db    *sql.DB
	ordersCache *map[string][]byte
}

func (o *orderService) AddOrder() error {

	return nil
}

func (o *orderService) GetOrderById(string) ([]byte, error) {

	return []byte("fd"), nil

}

func NewOrderService(db *sql.DB, cache *map[string][]byte) *orderService {
	return &orderService{db: db, ordersCache: cache}
}
