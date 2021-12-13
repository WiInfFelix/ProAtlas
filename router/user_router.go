package router

import (
	"ProAtlasV6/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitUserRouter(router *gin.Engine) {

	log.Println("Initializing user routes...")

	router.POST("/register", services.CreateUser)

	router.POST("/login", func(context *gin.Context) {
		context.String(http.StatusOK, "LoginRoute")
	})

	router.GET("/user/{user_id}", func(context *gin.Context) {
		context.String(http.StatusOK, "Get User by ID")
	})

	router.PUT("/user/{user_id}", func(context *gin.Context) {
		context.String(http.StatusOK, "Update user")
	})

	router.DELETE("/user/{user_id}", func(context *gin.Context) {
		context.String(http.StatusOK, "Delete User")
	})

	router.GET("/users", services.GetAllUsers)
}
