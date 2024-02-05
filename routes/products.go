package routes

import (
	"net/http"
	"strconv"

	"github.com/dpozzan/models"
	"github.com/gin-gonic/gin"
)

// var products []models.Product

func createProduct(context *gin.Context) {
	var product models.Product

	err := context.BindJSON(&product)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse body"})
		return
	}

	err = product.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func fetchProducts(context *gin.Context) {
	products, err := models.GetAllProducts()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch products"})
		return
	}

	context.JSON(http.StatusOK, products)
}

func fetchProduct(context *gin.Context) {
	product_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Wrong path"})
		return
	}

	product, err := models.GetProductByID(product_id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, *product)
}