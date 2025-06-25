// Package fibonacci provides functions for calculating Fibonacci numbers
package fibonacci

import (
	"fmt"
	"log"
	"math"
)

// Service represents a Fibonacci calculation service
type Service struct {
	logger *log.Logger
	cache  map[uint64]uint64
}

// NewService creates a new Fibonacci service
func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
		cache:  make(map[uint64]uint64),
	}
}

// Calculate returns the nth Fibonacci number using memoization
func (s *Service) Calculate(n uint64) (uint64, error) {
	s.logger.Printf("Calculating Fibonacci number for %v\n", n)
	
	// Check if the Fibonacci number is already cached
	if val, ok := s.cache[n]; ok {
		s.logger.Printf("Fibonacci number for %v is already memoised: %v\n", n, val)
		return val, nil
	}
	
	// Calculate the Fibonacci number
	var result uint64
	var val1 uint64
	if n <= 1 {
		result = n
	} else {
		val1, err := s.Calculate(n-1)
		if err != nil {
			return 0, err
		}
		
		val2, err := s.Calculate(n-2)
		if err != nil {
			return 0, err
		}
		
		result = val1 + val2
	}
	
	// Check for overflow
	if result > math.MaxUint64-val1 {
		return 0, fmt.Errorf("fibonacci number %d is too large and would cause overflow", n)
	}
	
	// Cache the result
	s.cache[n] = result
	s.logger.Printf("Fibonacci number for %v is %v\n", n, result)
	
	return result, nil
}

// GetCache returns a copy of the current cache
func (s *Service) GetCache() map[uint64]uint64 {
	cacheCopy := make(map[uint64]uint64)
	for k, v := range s.cache {
		cacheCopy[k] = v
	}
	return cacheCopy
}
