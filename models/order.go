package models

import "github.com/dpozzan/db"

type Order struct {
	ID         int64
	TotalPrice float64	`json:"total_price"`
	CustomerID int64	`json:"customer_id"`
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

	res, err := stmt.Exec(o.CustomerID, o.TotalPrice)

	if err != nil {
		return 0, err
	}

	order_id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return order_id, nil
}