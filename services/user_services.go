package services

import (
	"ProAtlasV6/database"
	"ProAtlasV6/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUsers(context *gin.Context) {

	users, err := database.GetAllUsers()

	if err != nil {
		context.JSON(500, err)
		return
	}

	context.JSON(http.StatusOK, users)

}

func CreateUser(context *gin.Context) {
	var NewUser models.User

	if err := context.BindJSON(&NewUser); err != nil {
		context.JSON(http.StatusBadRequest, "Provided invalid JSON")
		return
	}

	err := database.InsertUserIntoDB(NewUser)

	if err != nil {
		context.JSON(500, err)
		return
	}

}

func GetUserByID(context *gin.Context) {
	var DummyUser models.User
	userID, _ := strconv.Atoi(context.Param("userID"))

	DummyUser, err := database.GetUserByID(userID)

	if err != nil {
		context.JSON(500, err)
		return
	}

	context.JSON(200, DummyUser)

}
