package controllers

import (
	"github.com/IkhsanDzul/crud-golang/config"
	"github.com/IkhsanDzul/crud-golang/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(201, product)
}
