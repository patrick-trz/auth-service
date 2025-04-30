package main

import (
	"auth-service/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	r := gin.Default()

	r.Run(":8080")
}
