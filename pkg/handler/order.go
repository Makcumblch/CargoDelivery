package handler

import (
	"net/http"

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
	// projectId, err := getProjectId(c)
	// if err != nil {
	// 	return
	// }

	// cars, err := h.services.ICar.GetAllCars(projectId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, getAllCarsResponse{
	// 	Data: cars,
	// })
}

func (h *Handler) getOrderById(c *gin.Context) {
	// projectId, err := getProjectId(c)
	// if err != nil {
	// 	return
	// }

	// carId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid car id")
	// 	return
	// }

	// car, err := h.services.ICar.GetCarById(projectId, carId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, car)
}

func (h *Handler) updateOrder(c *gin.Context) {
	// projectId, err := getProjectId(c)
	// if err != nil {
	// 	return
	// }

	// carId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid car id")
	// 	return
	// }

	// var input cargodelivery.UpdateCar
	// if err := c.BindJSON(&input); err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// err = h.services.ICar.UpdateCar(projectId, carId, input)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, statusResponse{
	// 	Status: "ok",
	// })
}

func (h *Handler) deleteOrder(c *gin.Context) {
	// projectId, err := getProjectId(c)
	// if err != nil {
	// 	return
	// }

	// carId, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	newErrorResponse(c, http.StatusBadRequest, "invalid car id")
	// 	return
	// }

	// err = h.services.ICar.DeleteCar(projectId, carId)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// c.JSON(http.StatusOK, statusResponse{
	// 	Status: "ok",
	// })
}
