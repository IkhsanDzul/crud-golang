package controllers

import "github.com/gin-gonic/gin"

func CreateProduct(c *gin.Context) {
	// Logic to create a product
	c.JSON(201, gin.H{"message": "Product created"})
}