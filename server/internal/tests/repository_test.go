// counter_test.go
package tests

import (
	"counter-app/internal/repository"

	"reflect"
	"testing"
)

func TestInMemoryRepository(t *testing.T) {
    repo := repository.NewInMemoryRepository()

    // Test Create
    repo.Create("testCounter")
    value := repo.Get("testCounter");
    if value != 0 {
        t.Errorf("Create failed to create a new counter")
    }

    err:=repo.Create("testCounter2")
    if err !=nil {
        t.Errorf("Create failed")
    }

    // Test Increment
    repo.Increment("testCounter")
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
    expected := map[string]int{"testCounter": 1, "testCounter2": 0}
    if !reflect.DeepEqual(allCounters, expected) {
        t.Errorf("GetAll failed, got %v, want %v", allCounters, expected)
    }
}
