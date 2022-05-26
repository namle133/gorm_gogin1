package encode

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/domain"
	"net/http"
)

func ReadIndented(c *gin.Context, products []domain.Product) {
	c.IndentedJSON(http.StatusOK, products)
}

func ReadOneIndented(c *gin.Context, product *domain.Product) {
	c.IndentedJSON(http.StatusOK, product)
}

func DeleteIndented(c *gin.Context, id string) {
	c.String(http.StatusOK, "Deleted product with code = %s", id)
}

func UpdateIndented(c *gin.Context, id string) {
	c.String(http.StatusOK, "Updated product with code = %s", id)
}
