package routes

import (
	"net/http"

	"github.com/dpozzan/models"
	"github.com/gin-gonic/gin"
)

func createCartItem(context *gin.Context) {
	var cart_item models.CartItem

	err := context.BindJSON(&cart_item)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse body request"})
		return
	}

	carte_item_id, err := cart_item.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	cart_item.ID = carte_item_id
	context.JSON(http.StatusCreated, cart_item)
}