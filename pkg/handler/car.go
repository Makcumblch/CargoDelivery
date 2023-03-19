package handler

import (
	"net/http"
	"strconv"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCar(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	var car cargodelivery.Car
	if err := c.BindJSON(&car); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	carId, err := h.services.ICar.CreateCar(projectId, car)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": carId,
	})
}

type getAllCarsResponse struct {
	Data []cargodelivery.Car `json:"data"`
}

func (h *Handler) getAllCars(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	cars, err := h.services.ICar.GetAllCars(projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCarsResponse{
		Data: cars,
	})
}

func (h *Handler) getCarById(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	carId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid car id")
		return
	}

	car, err := h.services.ICar.GetCarById(projectId, carId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *Handler) updateCar(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	carId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid car id")
		return
	}

	var input cargodelivery.UpdateCar
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ICar.UpdateCar(projectId, carId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteCar(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	carId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid car id")
		return
	}

	err = h.services.ICar.DeleteCar(projectId, carId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
