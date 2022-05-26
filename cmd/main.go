package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/database"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/http/decode"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/http/encode"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/service"
	"net/http"
)

func main() {

	router := gin.Default()
	p := service.Product{
		Db: database.ConnectDatabase(),
	}
	var i service.IProduct = &p

	router.POST("/create", func(context *gin.Context) {
		it := decode.CreateBind(context)
		i.Create(it)
	})
	router.GET("/read", func(context *gin.Context) {
		products, err := i.Read()
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.ReadIndented(context, products)
	})
	router.GET("/read/:id", func(context *gin.Context) {
		id := decode.ReadOneBind(context)
		item, err := i.ReadOne(id)
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.ReadOneIndented(context, item)
	})
	router.DELETE("/delete/:id", func(context *gin.Context) {
		id := decode.DeleteOneBind(context)
		err := i.Delete(id)
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.DeleteIndented(context, id)

	})
	router.PUT("/update/:id", func(context *gin.Context) {
		id, price := decode.UpdateOneBind(context)
		fmt.Println(id, price)
		err := i.Update(id, price)
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.UpdateIndented(context, id)

	})

	router.Run(":8000")
}
