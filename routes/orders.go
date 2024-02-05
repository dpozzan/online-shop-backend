package routes

import (
	"net/http"

	"github.com/dpozzan/models"
	"github.com/gin-gonic/gin"
)

func createOrder(context *gin.Context) {
	var order models.Order

	err := context.BindJSON(&order)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H {"message": err.Error()})
		return
	}

	order_id, err := order.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	order.ID = order_id

	context.JSON(http.StatusCreated, order)
}