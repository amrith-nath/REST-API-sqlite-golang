package middlewares

import (
	"fmt"
	"net/http"

	"example.com/booking-app/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	fmt.Println(token)

	userId, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println(err)

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
		return
	}
	context.Set("user_id", userId)
	context.Next()
}
