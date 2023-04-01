package handler

import (
	"net/http"
	"strconv"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createOrder(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	var order cargodelivery.Order
	if err := c.BindJSON(&order); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	orderId, err := h.services.IOrder.CreateOrder(projectId, clientId, order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": orderId,
	})
}

type getAllOrdersResponse struct {
	Data []cargodelivery.Order `json:"data"`
}

func (h *Handler) getAllOrders(c *gin.Context) {
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	orders, err := h.services.IOrder.GetAllOrders(clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) getOrderById(c *gin.Context) {
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	orderId, err := strconv.Atoi(c.Param("idOrder"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid order id")
		return
	}

	order, err := h.services.IOrder.GetOrderById(clientId, orderId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) updateOrder(c *gin.Context) {
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	orderId, err := strconv.Atoi(c.Param("idOrder"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid order id")
		return
	}

	var input cargodelivery.UpdateOrder
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IOrder.UpdateOrder(clientId, orderId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteOrder(c *gin.Context) {
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	orderId, err := strconv.Atoi(c.Param("idOrder"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid order id")
		return
	}

	err = h.services.IOrder.DeleteOrder(clientId, orderId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
