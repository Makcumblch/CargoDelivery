package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	projectCtx          = "projectId"
	accessCtx           = "access"
	clientCtx           = "clientId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.IAuthorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getId(c *gin.Context, ctx, mesErrCtx, mesErrToInt string) (int, error) {
	id, ok := c.Get(ctx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, mesErrCtx)
		return 0, errors.New(mesErrCtx)
	}

	intId, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, mesErrToInt)
	}

	return intId, nil
}

func getUserId(c *gin.Context) (int, error) {
	return getId(c, userCtx, "user id not found", "invalid user id")
}

func (h *Handler) userAccessProject(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	projectId, err := strconv.Atoi(c.Param("idProject"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid project id")
		return
	}

	project, err := h.services.IProject.GetUserProject(userId, projectId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "project is not available")
		return
	}

	c.Set(projectCtx, projectId)
	c.Set(accessCtx, project.Access)
}

func getProjectId(c *gin.Context) (int, error) {
	return getId(c, projectCtx, "project id not found", "invalid project id")
}

func getAccess(c *gin.Context) (string, error) {
	id, ok := c.Get(accessCtx)
	if !ok {
		errorMes := "access not found"
		newErrorResponse(c, http.StatusInternalServerError, errorMes)
		return "", errors.New(errorMes)
	}

	strAccess, ok := id.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid access")
	}

	return strAccess, nil
}

func (h *Handler) accessOwner(c *gin.Context) {
	access, err := getAccess(c)
	if err != nil {
		return
	}

	if access != cargodelivery.OWNER {
		newErrorResponse(c, http.StatusForbidden, "missing rights owner")
		return
	}
}

func (h *Handler) accessWrite(c *gin.Context) {
	access, err := getAccess(c)
	if err != nil {
		return
	}

	if access != cargodelivery.OWNER && access != cargodelivery.WRITE {
		newErrorResponse(c, http.StatusForbidden, "missing rights owner/write")
		return
	}
}

func (h *Handler) userAccessClient(c *gin.Context) {
	projectId, err := getProjectId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	clientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid client id")
		return
	}

	_, err = h.services.IClient.GetClientById(projectId, clientId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client is not available")
		return
	}

	c.Set(clientCtx, clientId)
}

func getClientId(c *gin.Context) (int, error) {
	return getId(c, clientCtx, "client id not found", "invalid client id")
}
