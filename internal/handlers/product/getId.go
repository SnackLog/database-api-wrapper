package product

import (
	"log"
	"net/http"

	"github.com/SnackLog/database-api-wrapper/internal/database/product"
	"github.com/gin-gonic/gin"
)

func (p *ProductController) GetID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	product, err := product.GetProductByID(p.DB, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
