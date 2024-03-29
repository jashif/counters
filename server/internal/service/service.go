package service

import (
	"counter-app/internal/repository"
	"counter-app/models"
)

// CounterService defines the operations available on counters.
type CounterService interface {
	CreateCounter(name string)(error)
	IncrementCounter(name string)(int,error)
	GetCounterValue(name string) int
	GetAllCounters() []models.Counter
}

type service struct {
	repo repository.CounterRepository
}

func NewService(repo repository.CounterRepository) CounterService {
	return &service{repo: repo}
}

func (s *service) CreateCounter(name string)(error) {
	return s.repo.Create(name)
}

func (s *service) IncrementCounter(name string)(int,error) {
	return s.repo.Increment(name)
}

func (s *service) GetCounterValue(name string) int {
	return s.repo.Get(name)
}

func (s *service) GetAllCounters() []models.Counter {
	return s.repo.GetAll()
}
