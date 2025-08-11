package controllers

import (
	"github.com/IkhsanDzul/crud-golang/config"
	"github.com/IkhsanDzul/crud-golang/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, user)
}

func EditUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(204, nil)
}