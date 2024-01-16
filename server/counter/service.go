package counter

// CounterService defines the operations available on counters.
type CounterService interface {
	CreateCounter(name string)
	IncrementCounter(name string)
	GetCounterValue(name string) int
	GetAllCounters() map[string]int
}

type service struct {
	repo CounterRepository
}

func NewService(repo CounterRepository) CounterService {
	return &service{repo: repo}
}

func (s *service) CreateCounter(name string) {
	s.repo.Create(name)
}

func (s *service) IncrementCounter(name string) {
	s.repo.Increment(name)
}

func (s *service) GetCounterValue(name string) int {
	return s.repo.Get(name)
}

func (s *service) GetAllCounters() map[string]int {
	return s.repo.GetAll()
}
