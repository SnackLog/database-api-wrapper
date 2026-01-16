package health

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hc *HealthController) Get(c *gin.Context) {
	err := hc.DB.Ping(context.TODO(), nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Database connection is down"})
		return
	}
	c.Status(http.StatusOK)
}
