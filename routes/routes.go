package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/products", fetchProducts)
	server.POST("/products", createProduct)
	server.GET("/products/:id", fetchProduct)
	server.POST("/cart-items", createCartItem)
	server.POST("/orders", createOrder)
	server.POST("/signup", registerNewUser)
	server.POST("/login", loginUser)	
}

