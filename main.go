package main

import (
	"ProAtlasV6/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting application...")

	r := gin.Default()

	router.InitHelperRoutes(r)
	router.InitUserRouter(r)

	r.Run()
}
