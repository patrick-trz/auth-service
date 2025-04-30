package routes

import (
	"auth-service/internal/controllers"
	"auth-service/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUpAuthRoutes(r *gin.Engine) {
	r.POST("/register", middlewares.CustomerCreateMiddleware, controllers.Register)
	r.POST("/login", middlewares.CustomerLoginMiddleware, controllers.Login)
}
