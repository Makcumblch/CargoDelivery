package handler

import (
	"net/http"
	"strconv"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createClient(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	var client cargodelivery.Client
	if err := c.BindJSON(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	clientId, err := h.services.IClient.CreateClient(projectId, client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": clientId,
	})
}

type getAllClientsResponse struct {
	Data []cargodelivery.Client `json:"data"`
}

func (h *Handler) getAllClients(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	clients, err := h.services.IClient.GetAllClients(projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllClientsResponse{
		Data: clients,
	})
}

func (h *Handler) getClientById(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid client id")
		return
	}

	client, err := h.services.IClient.GetClientById(projectId, clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *Handler) updateClient(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid client id")
		return
	}

	var input cargodelivery.UpdateClient
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IClient.UpdateClient(projectId, clientId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteClient(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid client id")
		return
	}

	err = h.services.IClient.DeleteClient(projectId, clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
