// repository_test.go
package repository

import (
	"testing"

	"library-backend/models"

	"github.com/stretchr/testify/assert"
)

func TestAddPickUpSchedule(t *testing.T) {
	repo := NewBookRepository()

	schedule := models.PickUpSchedule{
		BookTitle:  "The Lord of the Rings",
		PickUpTime: "2023-10-20 10:00:00",
	}

	err := repo.AddPickUpSchedule(schedule)
	assert.NoError(t, err)

	storedSchedules := repo.GetAllPickUpSchedules()
	assert.Equal(t, 1, len(storedSchedules))
	assert.Equal(t, schedule, storedSchedules[0])
}

func TestGetAllPickUpSchedules(t *testing.T) {
	repo := NewBookRepository()

	schedule1 := models.PickUpSchedule{
		BookTitle:  "The Hobbit",
		PickUpTime: "2023-10-20 11:00:00",
	}

	schedule2 := models.PickUpSchedule{
		BookTitle:  "The Silmarillion",
		PickUpTime: "2023-10-21 12:00:00",
	}

	repo.AddPickUpSchedule(schedule1)
	repo.AddPickUpSchedule(schedule2)

	schedules := repo.GetAllPickUpSchedules()
	assert.Equal(t, 2, len(schedules))
	assert.Contains(t, schedules, schedule1)
	assert.Contains(t, schedules, schedule2)
}
