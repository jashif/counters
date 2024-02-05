package repository

import (
	"counter-app/models"
	"fmt"
	"sync"
)

// CounterRepository represents the data access layer.
type CounterRepository interface {
	Create(name string) error
	Increment(name string) (int, error)
	Get(name string) int
	GetAll() []models.Counter
}

type InMemoryRepository struct {
	counters []models.Counter
	mu       sync.Mutex
}

func NewInMemoryRepository() CounterRepository {
	return &InMemoryRepository{
		counters: []models.Counter{},
	}
}

func (r *InMemoryRepository) Create(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if name == "" {
		return fmt.Errorf("name is expected for the counter")
	}
	for _, c := range r.counters {
		if c.Name == name {
			return fmt.Errorf("key '%s' already exists", name)
		}
	}
	r.counters = append(r.counters, models.Counter{Name: name, Value: 0})
	return nil
}

func (r *InMemoryRepository) Increment(name string) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, c := range r.counters {
		if c.Name == name {
			r.counters[i].Value++
			return r.counters[i].Value, nil
		}
	}
	return 0, fmt.Errorf("key '%s' does not exist", name)
}

func (r *InMemoryRepository) Get(name string) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, c := range r.counters {
		if c.Name == name {
			return c.Value
		}
	}
	return 0
}

func (r *InMemoryRepository) GetAll() []models.Counter {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.counters
}
