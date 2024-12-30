package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/oganes5796/e-commerce/internal/models"
)

func CreateOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Проверяем наличие продукта и вычисляем общую цену
		var price float64
		err := db.QueryRow("SELECT price FROM products WHERE id = $1", order.ProductID).Scan(&price)
		if err != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		order.TotalPrice = price * float64(order.Quantity)

		// Сохраняем заказ в базу данных
		_, err = db.Exec("INSERT INTO orders (user_id, product_id, quantity, total_price) VALUES ($1, $2, $3, $4)",
			order.UserID, order.ProductID, order.Quantity, order.TotalPrice)
		if err != nil {
			http.Error(w, "Error creating order", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Order created successfully"})
	}
}

func GetOrdersByUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("user_id").(int)

		rows, err := db.Query("SELECT id, product_id, quantity, total_price FROM orders WHERE user_id = $1", userID)
		if err != nil {
			http.Error(w, "Error retrieving orders", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var orders []models.Order
		for rows.Next() {
			var order models.Order
			if err := rows.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.TotalPrice); err != nil {
				http.Error(w, "Error scanning order", http.StatusInternalServerError)
				return
			}
			orders = append(orders, order)
		}

		json.NewEncoder(w).Encode(orders)
	}
}
