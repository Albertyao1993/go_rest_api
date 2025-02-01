package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_rest_api/models"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. "})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user, try later"})
	}

	context.JSON(http.StatusCreated, gin.H{"message:": "Event created!", "user": user})
}
