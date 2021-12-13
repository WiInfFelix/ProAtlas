package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitHelperRoutes(r *gin.Engine) {

	log.Println("Initializing helper routes...")

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
}
