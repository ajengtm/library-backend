package handler

import (
	"encoding/json"
	"net/http"

	"library-backend/models"
	"library-backend/service"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{bookService}
}

func (h *BookHandler) GetBooksByGenre(w http.ResponseWriter, r *http.Request) {
	genre := r.URL.Query().Get("genre")
	if genre == "" {
		http.Error(w, "Genre is required", http.StatusBadRequest)
		return
	}

	books, err := h.bookService.GetBooksByGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":  "success",
		"message": "Books retrieved successfully",
		"data":    books,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *BookHandler) ScheduleBookPickUp(w http.ResponseWriter, r *http.Request) {
	var schedule models.PickUpSchedule
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := h.bookService.ScheduleBookPickUp(schedule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":  "success",
		"message": "Book pick-up scheduled successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *BookHandler) GetAllPickUpSchedules(w http.ResponseWriter, r *http.Request) {
	schedules := h.bookService.GetAllPickUpSchedules()
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":  "success",
		"message": "Pick-up schedules retrieved successfully",
		"data":    schedules,
	}
	json.NewEncoder(w).Encode(response)
}
