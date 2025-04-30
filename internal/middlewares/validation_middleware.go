package middlewares

import (
	"auth-service/internal/dtos"
	"auth-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CustomerCreateMiddleware(c *gin.Context) {
	var customer dtos.CustomerCreate
	if !utils.DecodeJSONRequest(c.Writer, c.Request, &customer) {
		return
	}

	if err := utils.ValidateStruct(customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Set("validatedCustomer", customer)
	c.Next()
}

func CustomerLoginMiddleware(c *gin.Context) {
	var customer dtos.CustomerLogin
	if !utils.DecodeJSONRequest(c.Writer, c.Request, &customer) {
		return
	}

	if err := utils.ValidateStruct(customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Set("validatedCustomer", customer)
	c.Next()
}
