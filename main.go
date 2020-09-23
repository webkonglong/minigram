package main

import (
	"log"
	
	"github.com/gin-gonic/gin"

	config "github.com/flyq/minigram/config"
	routes "github.com/flyq/minigram/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":80"))
}
