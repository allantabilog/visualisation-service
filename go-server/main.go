package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
<<<<<<< Updated upstream
	fibs = make(map[uint64]uint64)
=======
	fibs   = make(map[int]int)
>>>>>>> Stashed changes
)

type FibonacciResponse struct {
	Key   uint64 `json:"key"`
	Value uint64 `json:"value"`
	Result string `json:"result"`
}

func main() {
	// setup a simple http server
	const PORT = 8080

	dr := NewDebugRouter()
	dr.HandleFunc("/fibonacci", handleFibonacciRequest)
	dr.HandleFunc("/", handlePingRequest)
	dr.ListenAndServe()

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

	fib := FibonacciMemoised(uint64(n))
	logger.Printf(`Dump of fib cache: %v\n`, fibs)
	var result string
	if fib > 0 {
		result = "success"
	} else {
		result = "failure"
	}
	response := FibonacciResponse{Key: uint64(n), Value: fib, Result: result}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Returning response: %v\n", string(jsonResponse))
	w.Write(jsonResponse)
}

<<<<<<< Updated upstream
func FibonacciMemoised(n uint64) uint64 {
=======
// @todo: recursive fibonacci function is not efficient
// transform it to an interative function
func FibonacciMemoised(n int) int {
>>>>>>> Stashed changes
	logger.Printf("Calculating Fibonacci number for %v\n", n)
	// check if the Fibonacci number is already memoised
	if val, ok := fibs[n]; ok {
		logger.Printf("Fibonacci number for %v is already memoised: %v\n", n, val)
		return val
	}
	// calculate the Fibonacci number
	if n <= 1 {
		fibs[n] = n
	} else {
		fibs[n] = FibonacciMemoised(n-1) + FibonacciMemoised(n-2)
	}
	if fibs[n] > math.MaxInt64-1 {
		panic("Fibonacci number is too large")
	}
	logger.Printf("Fibonacci number for %v is %v\n", n, fibs[n])
	return fibs[n]
}
