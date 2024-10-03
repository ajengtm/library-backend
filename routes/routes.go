package routes

import (
	"library-backend/handler"

	"github.com/gorilla/mux"
)

func InitRoutes(bookHandler *handler.BookHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", bookHandler.GetBooksByGenre).Methods("GET")
	router.HandleFunc("/schedule", bookHandler.ScheduleBookPickUp).Methods("POST")
	router.HandleFunc("/schedules", bookHandler.GetAllPickUpSchedules).Methods("GET")

	return router
}
