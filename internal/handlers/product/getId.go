package product

import (
	"log"

	"github.com/SnackLog/database-api-wrapper/internal/database/product"
	"github.com/gin-gonic/gin"
)

func (p *ProductController) GetID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	productJSON, err := product.GetProductByID(p.DB, id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if productJSON == "" {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, gin.H{"product": productJSON})
}
