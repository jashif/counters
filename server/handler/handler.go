package handler

import (
	"counter-app/internal/service"
	"counter-app/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CounterHandler holds the dependencies for the counter-related HTTP handlers
type CounterHandler struct {
	service service.CounterService
}
type ErrorResponse struct {
    Error string `json:"error"`
}
func NewCounterHandler(s service.CounterService) *CounterHandler {
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

	keyExist := h.service.CreateCounter(req.Name)
	if keyExist != nil {
		// Create an error response
		errResp := ErrorResponse{
			Error: keyExist.Error(),
		}

		// Set the header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the status code
		w.WriteHeader(http.StatusBadRequest)

		// Encode the error response as JSON and send it
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			// Handle the error of failing to encode the response
			http.Error(w, "Failed to encode the error message", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CounterHandler) IncrementCounter(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	value,err := h.service.IncrementCounter(name)
	if err !=nil{
		errResp := ErrorResponse{
			Error: err.Error(),
		}

		w.WriteHeader(http.StatusNotFound)
		// Set the header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Encode the error response as JSON and send it
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			// Handle the error of failing to encode the response
			http.Error(w, "Failed to encode the error message", http.StatusInternalServerError)
		}
		return;
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{name: value})
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
