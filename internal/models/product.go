package models

import (
	"log"

	"github.com/oganes5796/e-commerce/internal/db"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

func (p *Product) Create() error {
	query := "INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := db.DB.QueryRow(query, p.Name, p.Description, p.Price, p.Stock).Scan(&p.ID); err != nil {
		log.Printf("Error creating product: %v", err)
		return err
	}
	return nil
}

func GetAllProducts() ([]Product, error) {
	rows, err := db.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Printf("Error getting all products: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Description, &p.Price, &p.Stock); err != nil {
			log.Printf("Error scanning product: %v", err)
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (p *Product) Update() error {
	query := "UPDATE products SET name = $1, description = $2, price = $3, stock = $4 WHERE id = $5"
	if _, err := db.DB.Exec(query, p.Name, p.Description, p.Price, p.Stock, p.ID); err != nil {
		log.Printf("Error updating product: %v", err)
		return err
	}
	return nil
}

func Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	if _, err := db.DB.Exec(query, id); err != nil {
		log.Printf("Error deleting product: %v", err)
		return err
	}

	return nil
}
