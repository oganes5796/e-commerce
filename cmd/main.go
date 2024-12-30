package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/oganes5796/e-commerce/internal/db"
	"github.com/oganes5796/e-commerce/internal/handlers"
)

func main() {
	// Заагрузка пеерменных окружения
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	// Подключение к БД
	db.Init()
	defer db.DB.Close()

	// Настройка маршрутов
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.RegistrUser(db.DB)).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser(db.DB)).Methods("POST")

	router.HandleFunc("/order", handlers.CreateOrder(db.DB)).Methods("POST")
	router.HandleFunc("/user/update", handlers.UpdateUser(db.DB)).Methods("PUT")
	router.HandleFunc("/user/delete", handlers.DeleteUser(db.DB)).Methods("DELETE")

	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Запуск сервера
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
