package services

import (
	"database/sql"
	"encoding/json"

	"github.com/m-a-r-a-t/L0/internal/http_server/models"
)

type orderRepo interface {
	GetOrderById(id string) (*models.Order, error)
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

	order, err := o.orderRepo.GetOrderById(id)

	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(order)

	if err != nil {
		return nil, err
	}

	return bytes, nil

}

func NewOrderService(db *sql.DB, cache *map[string][]byte, orderRepo orderRepo) *orderService {
	return &orderService{ordersCache: cache, orderRepo: orderRepo}
}
