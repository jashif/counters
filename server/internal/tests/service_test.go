// counter_test.go
package tests

import (
	"counter-app/internal/repository"
	"counter-app/internal/service"
	"testing"
)

func TestNewService(t *testing.T) {
    repo := repository.NewInMemoryRepository()
    newService:= service.NewService(repo);
    // Test Create
    newService.CreateCounter("testCounter")
    value := newService.GetCounterValue("testCounter");
    if value ==-1 {
        t.Errorf("Create failed to create a new counter")
    }
}
