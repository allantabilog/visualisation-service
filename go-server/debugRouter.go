package main

import (
	"fmt"
	"net/http"
)

const (
	defaultPort = 8080
)

type DebugRouter struct {
	mux    *http.ServeMux
	routes []string
}

func NewDebugRouter() *DebugRouter {
	return &DebugRouter{
		mux:    http.NewServeMux(),
		routes: []string{},
	}
}

func (dr *DebugRouter) HandleFunc(pattern string, handler http.HandlerFunc) {
	dr.routes = append(dr.routes, pattern)
	dr.mux.HandleFunc(pattern, handler)
}

func (dr *DebugRouter) LogRoutes() {
	for _, route := range dr.routes {
		logger.Printf("Registered route: %s", route)
	}
}
func (dr *DebugRouter) ServeHTTPOnPort(port int) {
	dr.LogRoutes()
	logger.Printf("DebugRouter is serving on port %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), dr.mux); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func (dr *DebugRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dr.mux.ServeHTTP(w, r)
}

func (dr *DebugRouter) ListenAndServe() {
	dr.ServeHTTPOnPort(defaultPort)
}
