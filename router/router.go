package router

import (
	"rest_api/controller"

	"github.com/gin-gonic/gin"
)

func StartApp(c controller.Repo) *gin.Engine {
	// Routes initialization
	r := gin.Default()

	// Create routes for activity
	v1 := r.Group("/activity-groups")
	{
		v1.GET("", c.GetActivity)
		v1.GET("/:id", c.GetActivityById)
		v1.POST("", c.AddActivity)
		v1.PATCH("/:id", c.UpdateActivity)
		v1.DELETE("/:id", c.DeleteActivityById)
	}

	// Create routes for todo items
	v2 := r.Group("/todo-items")
	{
		v2.GET("", c.GetTodo)
		v2.GET("/:id", c.GetTodoById)
		v2.POST("", c.AddTodo)
		v2.PATCH("/:id", c.UpdateTodo)
		v2.DELETE("/:id", c.DeleteTodoById)
	}
	return r
}
