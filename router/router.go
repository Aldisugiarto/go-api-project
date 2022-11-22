package router

import (
	"rest_api/controller"

	"github.com/gin-gonic/gin"
)

func StartApp(c controller.Repo) *gin.Engine {
	r := gin.Default()
	r.GET("/cakes", c.GetCake)
	r.GET("/cakes/:id", c.GetCakeById)
	r.POST("/cakes", c.AddCake)
	r.PATCH("/cakes/:id", c.UpdateCake)
	r.DELETE("/cakes/:id", c.DeleteCakeById)
	return r
}
