package service_test

import (
	"testing"

	"library-backend/models"
	"library-backend/repository"
	"library-backend/service"

	"github.com/stretchr/testify/assert"
)

func TestGetBooksByGenre(t *testing.T) {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)

	books, err := svc.GetBooksByGenre("science")

	assert.Nil(t, err)
	assert.NotEmpty(t, books)
}

func TestScheduleBookPickUp(t *testing.T) {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)

	schedule := models.PickUpSchedule{
		BookTitle:  "Test Book",
		PickUpTime: "2023-10-01T15:00:00Z",
	}

	err := svc.ScheduleBookPickUp(schedule)
	assert.Nil(t, err)
}
