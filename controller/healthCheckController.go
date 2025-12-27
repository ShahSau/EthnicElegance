package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary      Health Check
// @Description  Endpoint to check the health of the server
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /v1/ecommerce/health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Server is up and running"})
}
