package database

import "log"

func AutoMigrate() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	createCategoryTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL
	);`

	createBrandTable := `
	CREATE TABLE IF NOT EXISTS brands (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL
	);`

	createProductTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		type VARCHAR(20) NOT NULL,
		category_id INT,
		brand_id INT,
		FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
		FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE SET NULL
	);`

	createVariantTable := `
	CREATE TABLE IF NOT EXISTS product_variants (
		id SERIAL PRIMARY KEY,
		product_id INT NOT NULL,
		name VARCHAR(100) NOT NULL,
		price INT NOT NULL,
		stock INT NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	);`

	createOrderTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	createOrderItemTable := `
	CREATE TABLE IF NOT EXISTS order_items (
		id SERIAL PRIMARY KEY,
		order_id INT NOT NULL,
		variant_id INT NOT NULL,
		quantity INT NOT NULL,
		price INT NOT NULL,
		FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
		FOREIGN KEY (variant_id) REFERENCES product_variants(id) ON DELETE CASCADE
	);`

	queries := []string{
		createUserTable,
		createCategoryTable,
		createBrandTable,
		createProductTable,
		createVariantTable,
		createOrderTable,
		createOrderItemTable,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to run migration: %v\nQuery: %s", err, query)
		}
	}

	log.Println("Database migrated successfully")
}
