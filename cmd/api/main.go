package main

// @title Task Manager API
// @version 1.0
// @description This is a sample Task Manager API built in Go.
// @schemes https
// @host taskmanager-ekqd.onrender.com
// @BasePath /
//
// @securityDefinitions.apikey BearerAuth
// @name Authorization
// @in header
import (
	"fmt"

	_ "example.com/task_manager/docs"
	"example.com/task_manager/internal/database"
	"example.com/task_manager/internal/models"
	"example.com/task_manager/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Connect to database
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := gin.Default()

	// Enable CORS for Swagger
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	// Register routes
	routes.RegisterRoutes(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080")

	fmt.Println("Server is running...")
}
