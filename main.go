package main

import (
	"rest_api/config"
	"rest_api/controller"
	"rest_api/router"
)

func main() {
	db := config.InitDB()
	cakeRepo := controller.Repo{
		DB: db,
	}
	r := router.StartApp(cakeRepo)
	r.Run(":8080")
}
