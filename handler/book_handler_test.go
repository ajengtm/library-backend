package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"library-backend/handler"
	"library-backend/models"
	"library-backend/repository"
	"library-backend/service"

	"github.com/stretchr/testify/assert"
)

func TestGetBooksByGenre(t *testing.T) {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	hd := handler.NewBookHandler(svc)

	req, _ := http.NewRequest("GET", "/books?genre=science", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hd.GetBooksByGenre)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestScheduleBookPickUp(t *testing.T) {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	hd := handler.NewBookHandler(svc)

	schedule := models.PickUpSchedule{
		BookTitle:  "Test Book",
		PickUpTime: "2023-10-01T15:00:00Z",
	}

	body, _ := json.Marshal(schedule)
	req, _ := http.NewRequest("POST", "/schedule", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hd.ScheduleBookPickUp)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
