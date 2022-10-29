package http_server

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/m-a-r-a-t/L0/internal/http_server/controllers/order"
	"github.com/m-a-r-a-t/L0/internal/http_server/repositories"
	"github.com/m-a-r-a-t/L0/internal/http_server/services"
)

type httpServer struct {
	Server      *fiber.App
	ordersCache *map[string][]byte
	db          *sql.DB
}

func InitHttpServer(db *sql.DB, ordersCache *map[string][]byte) *httpServer {

	orderRepo := repositories.NewOrderRepo(db)

	orderService := services.NewOrderService(db, ordersCache, orderRepo)

	orderController := order.NewOrderController(orderService)

	server := fiber.New()
	server.Static("/", "./frontend")
	server.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders:  "Origin, Content-Type, Accept",
}))

	server.Get("/order", orderController.GetOrder)
	return &httpServer{
		Server:      server,
		ordersCache: ordersCache,
		db:          db,
	}

}
