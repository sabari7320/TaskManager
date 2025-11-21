package middleware

import (
	"fmt"
	"net/http"

	"example.com/task_manager/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized0.1."})
			context.Abort()
		}
		fmt.Println("authheader is", authHeader)
		// parts := strings.Split(authHeader, " ")
		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized 1.1."})
		// 	return
		// }

		token := authHeader

		userId, err := utils.VerifyToken(token)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized2."})
			return
		}
		context.Set("userId", userId)
		context.Next()

	}
}
