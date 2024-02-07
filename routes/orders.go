package routes

import (
	"net/http"
	"strconv"

	"github.com/dpozzan/models"
	"github.com/gin-gonic/gin"
)

type SetPriceOrderBody struct {
	Order models.Order `json:"order"`
	Cart_items []models.CartItem `json:"cart_items"`
}

func createOrder(context *gin.Context) {
	customer_id := context.GetInt64("userId")
	var order models.Order

	order.CustomerID = customer_id

	order_id, err := order.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	order.ID = order_id

	context.JSON(http.StatusCreated, order)
}

func getOrder(context *gin.Context) {
	order_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	user_id := context.GetInt64("userId")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse the ID from the URL"})
		return
	}

	order, err := models.GetOrderById(order_id, user_id)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, order)
}

func setOrderPrice(context *gin.Context){
	var orderBody SetPriceOrderBody
	user_id := context.GetInt64("userId")

	err := context.BindJSON(&orderBody)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse body"})
		return
	}

	order := orderBody.Order
	items := orderBody.Cart_items

	if order.CustomerID != user_id {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = order.SetPrice(items)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to set the total price of the order"})
		return
	}

	context.JSON(http.StatusOK, order)
}