package controllers

import (
	"jwt-golang/database"
	"jwt-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User

	//is sent by the client as a JSON body will be mapped into the user variable. It’s quite simple with gin.
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// is sent by the client as a JSON body will be mapped into the user variable. It’s quite simple with gin.
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// Once hashed, we store the user data into the database using the GORM global instance that we initialized earlier in the main file.
	record := database.Instance.Create(&user)

	// If there is an error while saving the data, the application would throw an HTTP Internal Server Error Code 500 and abort the request.
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	// Finally, if everything goes well, we send back the user id, name, and email to the client along with a 200 SUCCESS status code.
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
