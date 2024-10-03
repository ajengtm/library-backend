package main

import (
	"log"
	"net/http"

	"library-backend/handler"
	"library-backend/repository"
	"library-backend/routes"
	"library-backend/service"
)

func main() {
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	router := routes.InitRoutes(bookHandler)

	log.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
