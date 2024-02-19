package routes

import (
	"fmt"
	"net/http"

	"example.com/booking-app/models"
	"example.com/booking-app/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println("save error => ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create User. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created !"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "token generation failed"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}
