package routes

import (
	"example.com/task_manager/internal/handlers"
	"example.com/task_manager/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	//Register user
	r.POST("/signup", handlers.Register)
	//login user
	r.POST("/login", handlers.Login)

	authenticated := r.Group("/")
	authenticated.Use(middleware.Authenticate())
	{
		authenticated.POST("/tasks", handlers.CreateTask)
		authenticated.GET("/tasks", handlers.GetTasks)
		authenticated.PUT("/tasks/:id", handlers.UpdateTask)
		authenticated.DELETE("/tasks/:id", handlers.DeleteTask)
	}

}
