package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
	fibs = make(map[int]int)
)

// todo: use bigger-sized integers
type FibonacciResponse struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}

func main() {
	// setup a simple http server
	const PORT = 8080
	http.HandleFunc("/fibonacci", handleFibonacciRequest);
	http.HandleFunc("/", handlePingRequest);
	logger.Printf("Server is running on port %v\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
}

func handlePingRequest(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Request received from %v\n", r.RemoteAddr)
	var response map[string]string = make(map[string]string)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// add a message to the response map
	response["message"] = "fibonacci-server is up"
	response["timestamp"] = time.Now().Format(time.RFC3339)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Printf("Returning response: %v\n", string(jsonResponse))
	w.Write(jsonResponse)
}

func handleFibonacciRequest(w http.ResponseWriter, r *http.Request) {
	// handle fibonacci requests
	logger.Printf("Request received from %v\n", r.RemoteAddr)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	nStr := query.Get("n")
	if nStr == "" {
		http.Error(w, "Query parameter 'n' is required", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		http.Error(w, "Query parameter 'n' must be an integer", http.StatusBadRequest)
		return
	}

	result := FibonacciMemoised(n)
	logger.Printf(`Dump of fib cache: %v\n`, fibs)
	response := FibonacciResponse{Key: n, Value: result}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Returning response: %v\n", string(jsonResponse))
	w.Write(jsonResponse)
}

func FibonacciMemoised(n int) int {
	logger.Printf("Calculating Fibonacci number for %v\n", n)
	// check if the Fibonacci number is already memoised
	if val, ok := fibs[n]; ok {
		return val
	}
	// calculate the Fibonacci number
	if n <= 1 {
		fibs[n] = n
	} else {
		fibs[n] = FibonacciMemoised(n-1) + FibonacciMemoised(n-2)
	}
	if fibs[n] > math.MaxInt64 - 1 {
		panic("Fibonacci number is too large")
	}
	logger.Printf("Fibonacci number for %v is %v\n", n, fibs[n])
	return fibs[n]
}