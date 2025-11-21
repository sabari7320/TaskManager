package handlers

// Task Handlers
//
// This file contains handlers for task CRUD operations.
import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/task_manager/internal/database"
	"example.com/task_manager/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Summary      Create a new task
// @Description  Create a task for the authenticated user (JWT required)
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      models.Task  true  "Task data"
// @Success      200   {object}  models.Task
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /tasks [post]
// @Security     BearerAuth
func CreateTask(context *gin.Context) {

	userID := context.GetUint("userId")

	var task models.Task

	err := context.ShouldBindJSON(&task)
	fmt.Println("create -- task userID", userID)
	fmt.Println("error from create task", err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	task = models.Task{
		Title:  task.Title,
		Done:   task.Done,
		UserID: userID,
	}

	database.DB.Create(&task)
	context.JSON(http.StatusOK, task)
}

// GetTasks godoc
// @Summary      Get all tasks for authenticated user
// @Description  Returns tasks belonging to the logged-in user
// @Tags         tasks
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Task
// @Failure      401  {object}  map[string]string
// @Router       /tasks [get]
func GetTasks(context *gin.Context) {
	userID := context.GetUint("userId")

	var tasks []models.Task
	fmt.Println("Get -- task userID", userID)
	database.DB.Where("user_id = ?", userID).Find(&tasks)
	fmt.Println("Get -- task", tasks)
	context.JSON(http.StatusOK, tasks)
}

// UpdateTask godoc
// @Summary      Update a task
// @Description  Update a task by ID (only owner can update)
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Task ID"
// @Param        task  body      models.Task true  "Updated task data"
// @Success      200   {object}  models.Task
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /tasks/{id} [put]
// @Security     BearerAuth
func UpdateTask(c *gin.Context) {
	userID := c.GetUint("userId")

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	var task models.Task

	// Check if task exists
	if err := database.DB.First(&task, taskId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	// Check owner
	if task.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	database.DB.Model(&task).Updates(input)
	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Delete a task by ID (only owner can delete)
// @Tags         tasks
// @Produce      json
// @Param        id   path   int  true  "Task ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /tasks/{id} [delete]
// @Security     BearerAuth
func DeleteTask(c *gin.Context) {
	userID := c.GetUint("userId")

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	var task models.Task

	// Check if task exists
	if err := database.DB.First(&task, taskId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	// Check owner
	if task.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
