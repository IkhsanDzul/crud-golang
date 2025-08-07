package controllers

import (
	"strings"

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
	product.Name = strings.ToLower(product.Name)
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(201, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(200, products)
}

func GetProductByName(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Product name is required"})
		return
	}
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	product.Name = strings.ToLower(product.Name)
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(200, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Product ID is required"})
		return
	}
	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(204, nil)
}