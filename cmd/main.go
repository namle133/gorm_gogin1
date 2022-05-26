package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/http/decode"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/http/encode"
	"github.com/namle133/gorm_gogin1.git/gorm_gogin/service"
	"net/http"
)

func main() {

	router := gin.Default()
	var P service.IProduct

	router.POST("/create", func(context *gin.Context) {
		it := decode.CreateBind(context)
		fmt.Println("Hello, world!")
		fmt.Println(it)
		P.Create(it)
	})
	router.GET("/read", func(context *gin.Context) {
		products, err := P.Read()
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.ReadIndented(context, products)
	})
	router.GET("/read/:id", func(context *gin.Context) {
		id := decode.ReadOneBind(context)
		item, err := P.ReadOne(id)
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.ReadOneIndented(context, item)
	})
	router.DELETE("/delete/:id", func(context *gin.Context) {
		id := decode.DeleteOneBind(context)
		err := P.Delete(id)
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.DeleteIndented(context, id)

	})
	router.PUT("/update/:id", func(context *gin.Context) {
		id, price := decode.UpdateOneBind(context)
		err := P.Update(id, price)
		if err != nil {
			context.JSON(400, http.StatusBadRequest)
			return
		}
		encode.UpdateIndented(context, id)

	})

	router.Run(":8000")
}
