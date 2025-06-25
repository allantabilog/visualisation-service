# Makefile for Visualisation Service

# Variables
GO_SERVER_DIR := go-server
FRONTEND_DIR := .

.PHONY: all build build-backend build-frontend start start-backend start-frontend docker docker-backend docker-frontend clean help

# Default target
all: build start

# Build both backend and frontend
build: build-backend build-frontend

# Build backend
build-backend:
	@echo "Building Go backend..."
	cd $(GO_SERVER_DIR) && go build -o fibonacci-server

# Build frontend
build-frontend:
	@echo "Building React frontend..."
	npm install

# Start both backend and frontend
start: start-backend start-frontend

# Start the backend server
start-backend: build-backend
	@echo "Starting Go backend server..."
	cd $(GO_SERVER_DIR) && ./fibonacci-server &
	@echo "Backend server started at http://localhost:8080"

# Start the frontend server
start-frontend: build-frontend
	@echo "Starting React frontend..."
	npm start

# Docker build and run
docker: docker-backend docker-frontend

# Build and run backend with Docker
docker-backend:
	@echo "Building and running Go backend with Docker..."
	cd $(GO_SERVER_DIR) && docker build -t vis-service-go:latest .
	docker run -d -p 8080:8080 --name vis-service-go vis-service-go:latest
	@echo "Backend container started at http://localhost:8080"

# Build and run frontend with Docker
docker-frontend:
	@echo "Building and running React frontend with Docker..."
	docker build -t vis-service-node:latest .
	docker run -d -p 3000:3000 --name vis-service-node vis-service-node:latest
	@echo "Frontend container started at http://localhost:3000"

# Run tests
test: test-backend test-frontend

# Test backend
test-backend:
	@echo "Testing Go backend..."
	cd $(GO_SERVER_DIR) && go test -v ./...

# Test frontend
test-frontend:
	@echo "Testing React frontend..."
	npm test -- --watchAll=false --passWithNoTests

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	cd $(GO_SERVER_DIR) && rm -f fibonacci-server
	rm -rf node_modules build

# Stop Docker containers if running
docker-stop:
	@echo "Stopping Docker containers..."
	docker stop vis-service-go vis-service-node 2>/dev/null || true
	docker rm vis-service-go vis-service-node 2>/dev/null || true

# Help target
help:
	@echo "Visualisation Service Makefile"
	@echo "Usage:"
	@echo "  make build         - Build both backend and frontend"
	@echo "  make start         - Start both backend and frontend"
	@echo "  make docker        - Build and run with Docker"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make docker-stop   - Stop Docker containers"
	@echo "  make help          - Show this help message"
	@echo ""
	@echo "You can also use individual targets like:"
	@echo "  make build-backend, make start-frontend, etc."
