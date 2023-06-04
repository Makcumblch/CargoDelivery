package handler

import (
	"errors"
	"net/http"
	"strconv"

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
		if errors.Is(err, cargodelivery.ErrCreateRouteFewCars) ||
			errors.Is(err, cargodelivery.ErrCreateRouteCars) ||
			errors.Is(err, cargodelivery.ErrCreateRouteClient) ||
			errors.Is(err, cargodelivery.ErrCreateRouteDepo) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": routeId,
	})
}

type getAllRoutesResponse struct {
	Data []cargodelivery.RouteResponse `json:"data"`
}

func (h *Handler) getAllRoutes(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	routeResponse, err := h.services.IRoute.GetAllRoutes(projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllRoutesResponse{
		Data: routeResponse,
	})
}

// func (h *Handler) getRouteById(c *gin.Context) {
// 	projectId, err := getProjectId(c)
// 	if err != nil {
// 		return
// 	}

// 	carId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, "invalid route id")
// 		return
// 	}

// 	car, err := h.services.IRoute.GetRouteById(projectId, carId)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, car)
// }

func (h *Handler) deleteRoute(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	routeId, err := strconv.Atoi(c.Param("idRoute"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid route id")
		return
	}

	err = h.services.IRoute.DeleteRoute(projectId, routeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
