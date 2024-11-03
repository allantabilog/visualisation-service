package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var response map[string]string = make(map[string]string)

func main() {
	// setup a simple http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// add a message to the response map
		response["message"] = "json-server is up"
		response["timestamp"] = time.Now().Format(time.RFC3339)

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	})
	http.ListenAndServe(":8080", nil)
}