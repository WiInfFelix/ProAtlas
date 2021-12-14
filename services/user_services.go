package services

import (
	"ProAtlasV6/database"
	"ProAtlasV6/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	userID := context.Param("user_id")
	userIDint, err := strconv.Atoi(userID)
	fmt.Println(userIDint)

	DummyUser, err = database.GetUserByID(userIDint)

	if err != nil {
		context.JSON(404, err)
		return
	}

	context.JSON(200, DummyUser)

}

func DeleteUser(context *gin.Context) {
	userID, err := strconv.Atoi(context.Param("user_id"))
	if err != nil {
		log.Println("Provided invalid ID")
		context.JSON(http.StatusBadRequest, "Provided invalid ID")
	}

	err = database.DeleteUser(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "There was an error deleting the user")
	}

	context.JSON(http.StatusOK, fmt.Sprintf("Deleted user with id %d", userID))
}
