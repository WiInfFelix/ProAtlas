package main

import (
	"ProAtlasV6/database"
	"ProAtlasV6/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting application...")

	r := gin.Default()

	err := database.SetUpDB()
	if err != nil {
		log.Panicf("There was an error initializing the database:\n %s", err)
	}

	router.InitHelperRoutes(r)
	router.InitUserRouter(r)

	r.Run()
}
