package repository

import (
	"fmt"
	"sync"
)

// CounterRepository represents the data access layer.
type CounterRepository interface {
	Create(name string)(error)
	Increment(name string)(int,error)
	Get(name string) int
	GetAll() map[string]int
}

type InMemoryRepository struct {
	counters map[string]int
	mu       sync.Mutex
}

func NewInMemoryRepository() CounterRepository {
	return &InMemoryRepository{
		counters: make(map[string]int),
	}
}

func (r *InMemoryRepository) Create(name string)(error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if name == "" {
		return fmt.Errorf("name is expected for the counter")
	}
	 // Check if the key already exists
	 if _, exists := r.counters[name]; exists {
        // Key already exists, return an error
        return fmt.Errorf("key '%s' already exists", name)
    }

    // Key does not exist, create it with initial value 0
    r.counters[name] = 0
    return nil
}

func (r *InMemoryRepository) Increment(name string) (int, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    // Check if the key exists in the map
    if _, exists := r.counters[name]; !exists {
        // Return zero and an error if the key does not exist
        return 0, fmt.Errorf("key '%s' does not exist", name)
    }

    // If the key exists, increment the counter
    r.counters[name]++
    return r.counters[name], nil
}


func (r *InMemoryRepository) Get(name string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.counters[name]
}

func (r *InMemoryRepository) GetAll() map[string]int {
	return r.counters
}
