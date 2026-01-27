package health

import (
	"context"
	"net/http"

	"github.com/SnackLog/database-api-wrapper/internal/handlers"
	"github.com/gin-gonic/gin"
)

func (hc *HealthController) Get(c *gin.Context) {
	err := hc.DB.Ping(context.TODO(), nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, handlers.Error{Error: "Database connection is down"})
		return
	}
	c.Status(http.StatusOK)
}
