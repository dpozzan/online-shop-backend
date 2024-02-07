package routes

import (
	"github.com/dpozzan/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/products", fetchProducts)
	server.POST("/products", createProduct)
	server.GET("/products/:id", fetchProduct)
	server.POST("/cart-items", middlewares.Authenticate, createCartItem)
	server.POST("/orders", middlewares.Authenticate, createOrder)
	server.PUT("/orders", middlewares.Authenticate, setOrderPrice)
	server.GET("/orders/:id", middlewares.Authenticate, getOrder)
	server.POST("/signup", registerNewUser)
	server.POST("/login", loginUser)	
}

