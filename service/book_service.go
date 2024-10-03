package service

import (
	"library-backend/models"
	"library-backend/repository"
)

type BookService interface {
	GetBooksByGenre(genre string) ([]models.Book, error)
	ScheduleBookPickUp(schedule models.PickUpSchedule) error
	GetAllPickUpSchedules() []models.PickUpSchedule
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetBooksByGenre(genre string) ([]models.Book, error) {
	return s.repo.GetBooksByGenre(genre)
}

func (s *bookService) ScheduleBookPickUp(schedule models.PickUpSchedule) error {
	return s.repo.AddPickUpSchedule(schedule)
}

func (s *bookService) GetAllPickUpSchedules() []models.PickUpSchedule {
	return s.repo.GetAllPickUpSchedules()
}
