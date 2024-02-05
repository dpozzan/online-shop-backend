package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()

}

func createTables() {
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		price REAL NOT NULL,
		image_urls TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createProductsTable)
	if err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	createCartItemsTable := `
	CREATE TABLE IF NOT EXISTS cart_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		order_id INTEGER NOT NULL,
		FOREIGN KEY(product_id) REFERENCES products(id),
		FOREIGN KEY(order_id) REFERENCES orders(id)
	)
	`
	_, err = DB.Exec(createCartItemsTable)

	if err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	createOrdersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER NOT NULL,
		total_price REAL NOT NULL
	)
	`

	_, err = DB.Exec(createOrdersTable)

	if err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createUsersTable)

	if err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}
}