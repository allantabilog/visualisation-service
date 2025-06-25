package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/allantabilog/fibonacci-service/fibonacci"
)

// Response represents the standard API response structure
type Response struct {
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// FibonacciResponse represents the response for a Fibonacci number request
type FibonacciResponse struct {
	Key    uint64 `json:"key"`
	Value  uint64 `json:"value"`
	Result string `json:"result"`
}

// Handler contains all the HTTP handlers for the API
type Handler struct {
	logger  *log.Logger
	fibSvc  *fibonacci.Service
}

// NewHandler creates a new Handler
func NewHandler(logger *log.Logger, fibSvc *fibonacci.Service) *Handler {
	return &Handler{
		logger: logger,
		fibSvc: fibSvc,
	}
}

// HandlePing handles the / endpoint (ping)
func (h *Handler) HandlePing(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("Request received from %v\n", r.RemoteAddr)
	
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Prepare response
	response := Response{
		Message:   "fibonacci-server is up",
		Timestamp: time.Now(),
	}

	// Marshal response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	h.logger.Printf("Returning response: %v\n", string(jsonResponse))
	w.Write(jsonResponse)
}

// HandleFibonacci handles the /fibonacci endpoint
func (h *Handler) HandleFibonacci(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("Request received from %v\n", r.RemoteAddr)
	
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Parse query parameter
	query := r.URL.Query()
	nStr := query.Get("n")
	if nStr == "" {
		http.Error(w, "Query parameter 'n' is required", http.StatusBadRequest)
		return
	}

	// Convert parameter to integer
	n, err := strconv.ParseUint(nStr, 10, 64)
	if err != nil {
		http.Error(w, "Query parameter 'n' must be a valid non-negative integer", http.StatusBadRequest)
		return
	}

	// Calculate Fibonacci number
	fib, err := h.fibSvc.Calculate(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log cache state
	h.logger.Printf("Dump of fib cache: %v\n", h.fibSvc.GetCache())

	// Prepare response
	result := "success"
	response := FibonacciResponse{
		Key:    n,
		Value:  fib,
		Result: result,
	}

	// Marshal response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	h.logger.Printf("Returning response: %v\n", string(jsonResponse))
	w.Write(jsonResponse)
}
