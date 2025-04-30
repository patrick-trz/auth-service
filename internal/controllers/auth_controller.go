package controllers

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Register(c *gin.Context) {
	customerAny, exists := c.Get("validatedCustomer")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Invalid customer context",
		})
		return
	}
	var customer models.Customer
	if err := copier.Copy(&customer, &customerAny); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy data"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	customer.Password = string(hashedPassword)

	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.Customer
	customerAny, exists := c.Get("validatedCustomer")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Invalid customer context",
		})
		return
	}
	var customer models.Customer
	if err := copier.Copy(&customer, &customerAny); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy data"})
		return
	}

	config.DB.Where("email = ?", customer.Email).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(customer.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// JWT генерация
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
