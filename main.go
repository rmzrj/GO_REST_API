package main

import (
	"github.com/gin-gonic/gin"
	"rest_api_example.com/db"
	"rest_api_example.com/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

    routes.RegisterRoutes(server)

	server.Run(":8080")

}


