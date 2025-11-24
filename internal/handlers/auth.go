package handlers

// Auth Handlers
//
// This file contains user authentication handlers.
import (
	"fmt"
	"net/http"

	"example.com/task_manager/internal/database"
	"example.com/task_manager/internal/models"
	"example.com/task_manager/internal/requests"
	"example.com/task_manager/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary      Register a new user
// @Description  Create user with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      requests.UserRequest  true  "User data"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /signup [post]
func Register(context *gin.Context) {
	var req requests.UserRequest

	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}
	err = database.DB.Create(&user).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Register godoc
// @Summary      Register a new user
// @Description  Login user with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     requests.UserRequest  true  "User data"
// @Success      201   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /login [post]
func Login(context *gin.Context) {

	var req requests.UserRequest
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	var user models.User
	err = database.DB.Where("email = ?", req.Email).First(&user).Error

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		fmt.Print("database error")
		fmt.Print(err)
		return
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, req.Password)

	if !passwordIsValid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		fmt.Print("passwordhash error")
		fmt.Print(passwordIsValid)
		return
	}

	token, err := utils.GenerateToken(user.Email, int64(user.ID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "server error."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})

}
