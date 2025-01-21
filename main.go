package main

import (
	"fmt"
	"log"

	"github.com/Ayyasy123/dibimbing-take-home-test/config"
	"github.com/Ayyasy123/dibimbing-take-home-test/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.SetupUserRoutes(config.DB, r)
	routes.SetupEventRoutes(config.DB, r)
	routes.SetupTicketRoutes(config.DB, r)

	log.Println("Server running on port 8080")

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
