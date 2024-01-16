package handler

import (
	"counter-app/counter"
	"counter-app/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CounterHandler holds the dependencies for the counter-related HTTP handlers
type CounterHandler struct {
	service counter.CounterService
}

func NewCounterHandler(s counter.CounterService) *CounterHandler {
	return &CounterHandler{service: s}
}

func (h *CounterHandler) CreateCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateCounterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.service.CreateCounter(req.Name)
	w.WriteHeader(http.StatusCreated)
}

func (h *CounterHandler) IncrementCounter(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	h.service.IncrementCounter(name)
	w.WriteHeader(http.StatusOK)
}

func (h *CounterHandler) GetCounterValue(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	value := h.service.GetCounterValue(name)
	json.NewEncoder(w).Encode(map[string]int{name: value})
}

func (h *CounterHandler) GetAllCounters(w http.ResponseWriter, r *http.Request) {
	counters := h.service.GetAllCounters()
	json.NewEncoder(w).Encode(counters)
}

// RegisterRoutes sets up the routes for the counter application.
func RegisterRoutes(router *mux.Router, counterHandler *CounterHandler) {
	router.HandleFunc("/create", counterHandler.CreateCounter).Methods("POST")
	router.HandleFunc("/increment", counterHandler.IncrementCounter).Methods("GET")
	router.HandleFunc("/value", counterHandler.GetCounterValue).Methods("GET")
	router.HandleFunc("/counters", counterHandler.GetAllCounters).Methods("GET")
}
