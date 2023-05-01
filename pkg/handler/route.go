package handler

import (
	"net/http"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createRoute(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	var settingsRoute cargodelivery.RouteSettings
	if err := c.BindJSON(&settingsRoute); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	routeId, err := h.services.IRoute.CreateRoute(projectId, settingsRoute)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": routeId,
	})
}
