package controllers

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Logic for user login
	c.JSON(200, gin.H{"message": "User logged in"})
}

func Register(c *gin.Context) {
	// Logic for user registration
	c.JSON(201, gin.H{"message": "User registered"})
}
