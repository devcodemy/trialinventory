package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck() string {
	return "UP"
}

func Com() {
	rt := gin.Default()
	rt.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	rt.Run("localhost:8080")
}
