package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"library-backend/models"
)

type BookRepository interface {
	GetBooksByGenre(genre string) ([]models.Book, error)
	AddPickUpSchedule(schedule models.PickUpSchedule) error
	GetAllPickUpSchedules() []models.PickUpSchedule
}

type bookRepository struct {
	schedules []models.PickUpSchedule
}

func NewBookRepository() BookRepository {
	return &bookRepository{
		schedules: []models.PickUpSchedule{},
	}
}

func (r *bookRepository) GetBooksByGenre(genre string) ([]models.Book, error) {
	url := fmt.Sprintf("https://openlibrary.org/subjects/%s.json", genre)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Works []models.ResponseBook `json:"works"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	books := []models.Book{}
	for _, rb := range result.Works {
		books = append(books, models.Book{
			Title:         rb.Title,
			Author:        fmt.Sprintf("%s", rb.AuthorName),
			EditionNumber: fmt.Sprintf("%d", rb.EditionCount),
		})
	}

	return books, nil
}

func (r *bookRepository) AddPickUpSchedule(schedule models.PickUpSchedule) error {
	r.schedules = append(r.schedules, schedule)
	return nil
}

func (r *bookRepository) GetAllPickUpSchedules() []models.PickUpSchedule {
	return r.schedules
}
