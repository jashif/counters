package counter

import "sync"

// CounterRepository represents the data access layer.
type CounterRepository interface {
	Create(name string)
	Increment(name string)
	Get(name string) int
	GetAll() map[string]int
}

type inMemoryRepository struct {
	counters map[string]int
	mu       sync.Mutex
}

func NewInMemoryRepository() CounterRepository {
	return &inMemoryRepository{
		counters: make(map[string]int),
	}
}

func (r *inMemoryRepository) Create(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counters[name] = 0
}

func (r *inMemoryRepository) Increment(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counters[name]++
}

func (r *inMemoryRepository) Get(name string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.counters[name]
}

func (r *inMemoryRepository) GetAll() map[string]int {
	return r.counters
}
