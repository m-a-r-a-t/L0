package order

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/m-a-r-a-t/L0/pkg/validator"
)

type getOrderController struct {
	orderService IorderService
}

func (c *getOrderController) GetOrder(f *fiber.Ctx) error {
	// queryData := validator.ValidatQueryeAndGetData[GetOrderQuery](f.QueryParser())
	order, err := c.orderService.GetOrderById("fdsfdsf")
	if err != nil {
		return err
	}
	return f.Send(order)

}

func NewOrderController(orderService IorderService) *getOrderController {
	return &getOrderController{orderService: orderService}
}
