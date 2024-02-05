package models

import (
	"encoding/json"

	"github.com/dpozzan/db"
)

type Product struct {
	ID          int64 `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	ImageUrls []string `json:"image_urls" binding:"required"`
}

func (p *Product) Save() error {
	query := `
	INSERT INTO products (name, description, price, image_urls)
	VALUES (?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// Serialize ImageUrls slice to JSON
	imageUrlsJson, err := json.Marshal(p.ImageUrls)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(p.Name, p.Description, p.Price, string(imageUrlsJson))

	if err != nil {
		return err
	}

	return nil
}


func GetAllProducts() ([]Product, error) {
	query := "SELECT * FROM products"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}


	var products []Product
	for rows.Next() {
		var product Product
		var imageUrlsJson string

		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &imageUrlsJson)

		if err != nil {
			return nil, err
		}

		// Deserialize image URLs from JSON
		err = json.Unmarshal([]byte(imageUrlsJson), &product.ImageUrls)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetProductByID(product_id int64) (*Product, error) {
	var product Product
	var imageUrlsJson string

	query := "SELECT * FROM products WHERE id = ?"
	

	row := db.DB.QueryRow(query, product_id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &imageUrlsJson)

	if err != nil {
		return nil, err
	}

	// Deserialize image URLs from JSON
	err = json.Unmarshal([]byte(imageUrlsJson), &product.ImageUrls)
	if err != nil {
		return nil, err
	}

	return &product, nil
}