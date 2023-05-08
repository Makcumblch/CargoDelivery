package handler

import (
	"net/http"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDepo(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	var depo cargodelivery.Depo
	if err := c.BindJSON(&depo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	depoId, err := h.services.IDepo.CreateDepo(projectId, depo)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": depoId,
	})
}

type getDepoResponse struct {
	Data cargodelivery.Client `json:"data"`
}

func (h *Handler) getDepo(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	depo, err := h.services.IDepo.GetDepo(projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getDepoResponse{
		Data: depo,
	})
}

func (h *Handler) updateDepo(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	var input cargodelivery.UpdateDepo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IDepo.UpdateDepo(projectId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
