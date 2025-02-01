package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go_rest_api/db"
	"github.com/go_rest_api/routes"
)

func main() {
	fmt.Println("Hello word")

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(": 8080") // localhost:8080
}
