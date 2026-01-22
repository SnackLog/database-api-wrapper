package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/database-api-wrapper/internal/database/product"
	"github.com/gin-gonic/gin"
)

// Get godoc
// @Summary      Search products
// @Description  Search for products by name
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        q      query     string  true  "Search query (min 3 chars)"
// @Param        limit  query     int     false "Limit results (default 10, max 100)"
// @Success      200  {object}  map[string][]product.Product
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     Bearer
// @Router       /products/search [get]
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

	products, err := product.SearchProductByName(p.DB, query, int(limit))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if products == nil {
		log.Println("Products was nil?!")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "products was nil."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
