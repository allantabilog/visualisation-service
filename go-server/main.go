package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
)

func main() {
	// setup a simple http server
	const PORT = 8080
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