package main

import (
	"github.com/dpozzan/db"
	"github.com/dpozzan/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	db.InitDB()

	server := gin.Default()
	corsConfig := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Headers you want to allow
		AllowCredentials: true,
	})
	server.Use(func(c *gin.Context) {
		corsConfig.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})
	server.Static("/images", "./images")
	
	routes.RegisterRoutes(server)
	server.Run(":8080")
}