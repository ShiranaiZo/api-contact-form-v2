package handlers

import (
	"net/http"

	"github.com/ShiranaiZo/api-contact-form-v2/responses"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "Success",
		Message: "API is runing.",
	})
}
