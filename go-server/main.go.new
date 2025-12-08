package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/allantabilog/fibonacci-service/fibonacci"
	"github.com/allantabilog/fibonacci-service/handlers"
)

func main() {
	// Setup logging to both file and stdout
	logDir := "logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	
	// Create log file with date in filename
	timestamp := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, fmt.Sprintf("fibonacci-server-%s.log", timestamp))
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	// Don't close the file as it needs to stay open for logging
	// defer file.Close()
	
	// Create a multi-writer to log to both file and stdout
	multiWriter := io.MultiWriter(file, os.Stdout)
	logger := log.New(multiWriter, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	
	logger.Printf("Logging initialized. Log file: %s", logFile)
	
	// Create services and handlers
	fibService := fibonacci.NewService(logger)
	apiHandler := handlers.NewHandler(logger, fibService)
	
	// setup a simple http server
	const PORT = 8080

	dr := NewDebugRouter()
	dr.HandleFunc("/fibonacci", apiHandler.HandleFibonacci)
	dr.HandleFunc("/", apiHandler.HandlePing)
	dr.ListenAndServe()
}
