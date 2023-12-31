package main

import (
	"github.com/gin-gonic/gin"

	"go-course/api/db"
	"go-course/api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
