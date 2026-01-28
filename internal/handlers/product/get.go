package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/database-api-wrapper/internal/database/product"
	"github.com/SnackLog/database-api-wrapper/internal/handlers"
	"github.com/gin-gonic/gin"
)

type productGetResponse struct {
	Products []product.Product `json:"products"`
}

// Get godoc
// @Summary      Search products
// @Description  Search for products by name
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        q      query     string  true  "Search query (min 3 chars)"
// @Param        limit  query     int     false "Limit results (default 10, max 100)"
// @Success      200  {object}  productGetResponse
// @Failure      400  {object}  handlers.Error
// @Failure      401  {object}  handlers.Error
// @Failure      500  {object}  handlers.Error
// @Security     Bearer
// @Router       /products/search [get]
func (p *ProductController) Get(c *gin.Context) {
	query := c.Query("q")
	if query == "" || len(query) < 3 {
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Query parameter 'q' is required and must be at least 3 characters long"})
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
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Query parameter 'limit' must be a positive integer less than 100"})
		return
	}

	products, err := product.SearchProductByName(p.DB, query, int(limit))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: err.Error()})
		return
	}
	if products == nil {
		log.Println("Products was nil?!")
		c.AbortWithStatusJSON(http.StatusInternalServerError, handlers.Error{Error: "products was nil."})
		return
	}

	c.JSON(http.StatusOK, productGetResponse{Products: *products})
}
