package services

import (
	"ProAtlasV6/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var testUsers = []models.User{
	{Id: 1, Username: "Lecker"},
	{Id: 2, Username: "Maes"},
}

func GetAllUsers(context *gin.Context) {

	context.JSON(http.StatusOK, testUsers)

}

func CreateUser(context *gin.Context) {
	var NewUser models.User

	if err := context.BindJSON(&NewUser); err != nil {
		context.JSON(http.StatusBadRequest, "Provided invalid JSON")
	}

	testUsers = append(testUsers, NewUser)

}
