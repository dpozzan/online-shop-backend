package models

import "github.com/dpozzan/db"

type CartItem struct {
	ID        int64	`json:"id"`
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int64 `json:"quantity" binding:"required"`
	OrderID int64 `json:"order_id" binding:"required"`
}

// var shopping_cart []CartItem

func (c *CartItem) Save() (int64, error) {
	query := `INSERT INTO cart_items (product_id, quantity, order_id)
	VALUES (?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(c.ProductID, c.Quantity, c.OrderID)

	if err != nil {
		return 0, err
	}

	cart_item_id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return cart_item_id, nil
}