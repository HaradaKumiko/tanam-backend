package admin

import (
	"net/http"
	"tanam-backend/helpers/response"
	"tanam-backend/services/admin"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	service admin.OrderService
}

func InitOrderController() OrderController {
	return OrderController{
		service: admin.InitOrderService(),
	}
}

func (controller *OrderController) GetAllOrderController(c echo.Context) error {
	orders, err := controller.service.GetAllOrderService()

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	if len(orders) < 1 {
		response := response.ErrorFormatter("Data Order Tidak Ditemukan")
		return c.JSON(http.StatusOK, response)
	}

	response := response.SuccessPluralFormatter("Data Semua Order", orders)
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderController) GetOrderByOrderIdController(c echo.Context) error {
	orderId := c.Param("orderId")
	order, err := controller.service.GetOrderByIdService(orderId)

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusNotFound, response)
	}

	response := response.SuccessSingularFormatter("Data Order", order)
	return c.JSON(http.StatusOK, response)
}
