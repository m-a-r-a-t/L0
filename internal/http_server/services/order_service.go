package services

import (
	"database/sql"
)

type orderRepo interface {
	GetOrderById(id string)
}

type orderService struct {
	ordersCache *map[string][]byte
	orderRepo   orderRepo
}

func (o *orderService) AddOrder() error {

	return nil
}

func (o *orderService) GetOrderById(id string) ([]byte, error) {
	if data, ok := (*o.ordersCache)[id]; ok {
		return data, nil
	}

	return []byte("fd"), nil

}

func NewOrderService(db *sql.DB, cache *map[string][]byte, orderRepo orderRepo) *orderService {
	return &orderService{ordersCache: cache, orderRepo: orderRepo}
}
