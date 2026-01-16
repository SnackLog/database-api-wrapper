package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/database-api-wrapper/internal/database/product"
	"github.com/gin-gonic/gin"
)

func (p *ProductController) Get(c *gin.Context) {
	query := c.Query("q")
	if query == "" || len(query) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required and must be at least 3 characters long"})
		return
	}

	limitString := c.Query("limit")
	if limitString == "" {
		limitString = "10"
	}

	limit, err := strconv.ParseInt(limitString, 10, 32)
	if err != nil || limit <= 0 || limit >= 100 {
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'limit' must be a positive integer less than 100"})
		return
	}

	productsJSON, err := product.SearchProductByName(p.DB, query, int(limit))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": productsJSON})
}
