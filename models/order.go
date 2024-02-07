package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dpozzan/db"
)

type Order struct {
	ID         int64	`json:"id"`
	TotalPrice float64	`json:"total_price"`
	CustomerID int64	`json:"customer_id"`
}

type ProductPrice struct {
	ID int64
	Price float64
	Quantity int64
}

func (o *Order) Save() (int64, error) {
	query := `
	INSERT INTO orders (customer_id, total_price)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(o.CustomerID, 0)

	if err != nil {
		return 0, err
	}

	order_id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return order_id, nil
}

func GetOrderById(order_id, user_id int64) (Order, error) {
	var order Order
	query := "SELECT * FROM orders WHERE id = ?"

	row := db.DB.QueryRow(query, order_id)

	err := row.Scan(&order.ID, &order.TotalPrice, &order.CustomerID)

	if err != nil {
		return Order{}, err
	}

	if order.CustomerID != user_id {
		return Order{}, errors.New("unauthorized")
	}

	return order, nil
}

func (o *Order) SetPrice(items []CartItem) (error) {

	total_price := 0.00

	product_quantities := make(map[int64]int64)

	for _, item := range items {
		product_quantities[item.ProductID] += item.Quantity
	}

	var product_ids []int64

	for id := range product_quantities {
		product_ids = append(product_ids, id)
	}

	placeholders := make([]string, len(product_ids))

	for i := range placeholders {
		placeholders[i] = "?"
	}

	query := fmt.Sprintf("SELECT id, price FROM products WHERE id IN (%s)", strings.Join(placeholders, ","))

	args := make([]interface{}, len(product_ids))

	for i, id := range product_ids {
		args[i] = id
	}

	rows, err := db.DB.Query(query, args...)

	if err != nil {
		return err
	}

	defer rows.Close()

	var product_prices []ProductPrice

	for rows.Next() {
		var pp ProductPrice

		err := rows.Scan(&pp.ID, &pp.Price)

		if err != nil {
			return err
		}

		pp.Quantity = product_quantities[pp.ID]
		product_prices = append(product_prices, pp)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	for _, product := range product_prices {
		total_price += float64(product.Quantity) * product.Price
	}

	query_price := "UPDATE orders SET total_price = ? WHERE id = ?"

	stmt, err := db.DB.Prepare(query_price)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(total_price, o.ID)

	if err != nil {
		return err
	}

	o.TotalPrice = total_price

	return nil

}