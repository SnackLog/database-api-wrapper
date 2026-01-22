package product

import (
	"log"
	"net/http"

	"github.com/SnackLog/database-api-wrapper/internal/database/product"
	"github.com/gin-gonic/gin"
)

// GetID godoc
// @Summary      Get product by ID
// @Description  Get a single product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  map[string]product.Product
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     Bearer
// @Router       /products/{id} [get]
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
