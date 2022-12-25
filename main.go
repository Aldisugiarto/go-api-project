package main

import (
	"rest_api/config"
	"rest_api/controller"
	"rest_api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB(&gin.Context{})
	cakeRepo := controller.Repo{
		DB: db,
	}
	r := router.StartApp(cakeRepo)
	r.Run(":3030")
}
