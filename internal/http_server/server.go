package http_server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/m-a-r-a-t/L0/internal/http_server/order"
	"github.com/m-a-r-a-t/L0/internal/services"
)

type httpServer struct {
	Server      *fiber.App
	ordersCache *map[string][]byte
	db          *sql.DB
}

func InitHttpServer(db *sql.DB, ordersCache *map[string][]byte) *httpServer {

	orderService := services.NewOrderService(db, ordersCache)

	orderController := order.NewOrderController(orderService)

	server := fiber.New()

	server.Get("/order:id", orderController.GetOrder)

	return &httpServer{
		Server:      server,
		ordersCache: ordersCache,
		db:          db,
	}

}
