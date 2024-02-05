// counter_test.go
package tests

import (
	"counter-app/internal/repository"
	"counter-app/models"
	"testing"
)

func TestInMemoryRepository(t *testing.T) {
	repo := repository.NewInMemoryRepository()

	// Test Create
	err := repo.Create("testCounter")
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	value := repo.Get("testCounter")
	if value != 0 {
		t.Errorf("Create failed to create a new counter, got %d", value)
	}

	err = repo.Create("testCounter2")
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Test Increment
	_, err = repo.Increment("testCounter")
	if err != nil {
		t.Errorf("Increment failed: %v", err)
	}

	if count := repo.Get("testCounter"); count != 1 {
		t.Errorf("Increment failed, got %d, want %d", count, 1)
	}

	// Test Get
	count := repo.Get("testCounter")
	if count != 1 {
		t.Errorf("Get failed, got %d, want %d", count, 1)
	}

	// Test GetAll
	allCounters := repo.GetAll()
	expected := []models.Counter{{Name: "testCounter", Value: 1}, {Name: "testCounter2", Value: 0}}

	for _, expCounter := range expected {
		found := false
		for _, counter := range allCounters {
			if counter.Name == expCounter.Name && counter.Value == expCounter.Value {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("GetAll failed, expected counter %v not found in %v", expCounter, allCounters)
		}
	}
}
