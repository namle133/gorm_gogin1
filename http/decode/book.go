package decode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/domain"
	"strconv"
)

func CreateBind(c *gin.Context) *domain.Product {
	var item *domain.Product
	err := c.BindJSON(&item)
	if err != nil {
		c.Error(err)
	}
	return item
}

func ReadOneBind(c *gin.Context) string {
	code := c.Param("id")
	return code
}

func DeleteOneBind(c *gin.Context) string {
	code := c.Param("id")
	return code
}

func UpdateOneBind(c *gin.Context) (string, int) {
	code := c.Param("id")
	pstr := c.Query("price")
	fmt.Println(code, pstr)
	price, err := strconv.Atoi(pstr)
	if err != nil {
		c.Error(err)
		return "Error", 0
	}
	return code, price
}
