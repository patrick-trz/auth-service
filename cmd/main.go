package main

import (
	"auth-service/internal/config"
	routes "auth-service/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	r := gin.Default()
	routes.SetUpAuthRoutes(r)

	r.Run(":8081")
}
