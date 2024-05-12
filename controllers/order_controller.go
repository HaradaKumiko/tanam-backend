package controllers

import (
	"net/http"
	"tanam-backend/domains/web/order"
	"tanam-backend/helpers/response"
	"tanam-backend/services"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	service services.OrderService
}

func InitOrderController() OrderController {
	return OrderController{
		service: services.InitOrderService(),
	}
}

func (controller *OrderController) CreateOrderController(c echo.Context) error {

	authId := c.Get("authId").(string)

	var order order.CreateOrderRequest
	if err := c.Bind(&order); err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := controller.service.CreateOrderService(order, authId)

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := response.SuccessSingularFormatter("Berhasil Order", data)
	return c.JSON(http.StatusOK, response)

}

func (controller *OrderController) GetOrderDonorController(c echo.Context) error {

	authId := c.Get("authId").(string)

	order, err := controller.service.GetOrderDonorService(authId)

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := response.SuccessPluralFormatter("Data Semua Order Pengguna", order)
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