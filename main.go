package main

import (
	"github.com/IkhsanDzul/crud-golang/config"
	"github.com/IkhsanDzul/crud-golang/controllers"
	"github.com/IkhsanDzul/crud-golang/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	//Connect to database
	config.ConnectDatabase()
	config.Migrate()

	r := gin.Default()
	
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.Use(middleware.AuthRequired())
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}

		master := api.Group("/master")
		{
			// master.GET("/products", controllers.GetProducts)
			// master.GET("/products/:id", controllers.GetProduct)
			master.POST("/products", controllers.CreateProduct)
			// master.PUT("/products/:id", controllers.UpdateProduct)
			// master.DELETE("/products/:id", controllers.DeleteProduct)
		}
	}

	r.Run(":9000")
}