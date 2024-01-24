package main

import (
	"counter-app/handler"
	"counter-app/internal/repository"
	"counter-app/internal/service"
	"counter-app/logger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewInMemoryRepository()
	service := service.NewService(repo)
	counterHandler := handler.NewCounterHandler(service)

	// Initialize the mux router
	r := mux.NewRouter()
	// Register the routes using the new function in handler package
	handler.RegisterRoutes(r, counterHandler)

	// Start the server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", CORS(logger.NewLogger(r))))
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
