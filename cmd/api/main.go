// @title Task Manager API
// @version 1.0
// @description This is a sample Task Manager API built in Go.
// @host localhost:8080
// @BasePath /
//
// @securityDefinitions.apikey BearerAuth
// @name Authorization
// @in header

package main

import (
	_ "example.com/task_manager/internal/handlers"

	"fmt"

	"example.com/task_manager/internal/database"
	"example.com/task_manager/internal/models"
	"example.com/task_manager/internal/routes"

	_ "example.com/task_manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin swagger middleware
)

func main() {
	// Connect to database
	database.Connect()

	// Migrate your models
	database.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// 👉 Add Swagger endpoint just before running server
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")

	fmt.Println("Hello, World!")
}
