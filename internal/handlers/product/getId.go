package product

import (
	"log"
	"net/http"

	"github.com/SnackLog/database-api-wrapper/internal/database/product"
	"github.com/SnackLog/database-api-wrapper/internal/handlers"
	"github.com/gin-gonic/gin"
)

type getProductByIdResponse struct {
	Product product.Product `json:"product"`
}

// GetID godoc
// @Summary      Get product by ID
// @Description  Get a single product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  map[string]product.Product
// @Failure      400  {object}  handlers.Error
// @Failure      401  {object}  handlers.Error
// @Failure      404  {object}  handlers.Error
// @Failure      500  {object}  handlers.Error
// @Security     Bearer
// @Router       /products/{id} [get]
func (p *ProductController) GetID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "ID is required"})
		return
	}

	product, err := product.GetProductByID(p.DB, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: err.Error()})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, handlers.Error{Error: "Product not found"})
		return
	}

	c.JSON(http.StatusOK, getProductByIdResponse{Product: *product})
}
