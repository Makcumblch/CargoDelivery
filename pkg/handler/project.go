package handler

import (
	"net/http"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createProject(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var project cargodelivery.Project
	if err := c.BindJSON(&project); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	projectId, err := h.services.IProject.CreateProject(userId, project)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": projectId,
	})
}

type getAllProjectsResponse struct {
	Data []cargodelivery.Project `json:"data"`
}

func (h *Handler) getAllProjects(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	projects, err := h.services.IProject.GetAllProjects(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllProjectsResponse{
		Data: projects,
	})
}

func (h *Handler) getProjectById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	project, err := h.services.IProject.GetProjectById(userId, projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *Handler) updateProject(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	var input cargodelivery.UpdateProject
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.IProject.UpdateProject(userId, projectId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteProject(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	projectId, err := getProjectId(c)
	if err != nil {
		return
	}

	err = h.services.IProject.DeleteProject(userId, projectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
