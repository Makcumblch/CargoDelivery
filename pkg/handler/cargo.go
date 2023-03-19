package handler

import (
	"net/http"
	"strconv"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCargo(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	var cargo cargodelivery.Cargo
	if err := c.BindJSON(&cargo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cargoId, err := h.services.ICargo.CreateCargo(projectId, cargo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": cargoId,
	})
}

type getAllCargosResponse struct {
	Data []cargodelivery.Cargo `json:"data"`
}

func (h *Handler) getAllCargos(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	cars, err := h.services.ICargo.GetAllCargos(projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCargosResponse{
		Data: cars,
	})
}

func (h *Handler) getCargoById(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	cargoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cargo id")
		return
	}

	cargo, err := h.services.ICargo.GetCargoById(projectId, cargoId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cargo)
}

func (h *Handler) updateCargo(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	cargoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cargo id")
		return
	}

	var input cargodelivery.UpdateCargo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ICargo.UpdateCargo(projectId, cargoId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteCargo(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	cargoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cargo id")
		return
	}

	err = h.services.ICargo.DeleteCargo(projectId, cargoId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
