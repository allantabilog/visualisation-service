package fibonacci

import (
	"io"
	"log"
	"testing"
)

func TestFibonacciCalculate(t *testing.T) {
	// Create a logger that discards output to keep tests quiet
	logger := log.New(io.Discard, "", 0)
	service := NewService(logger)

	testCases := []struct {
		n        uint64
		expected uint64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{10, 55},
		{20, 6765},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result, err := service.Calculate(tc.n)
			if err != nil {
				t.Errorf("Expected no error for n=%d, got %v", tc.n, err)
			}
			if result != tc.expected {
				t.Errorf("Calculate(%d) = %d, want %d", tc.n, result, tc.expected)
			}
		})
	}
}

func TestFibonacciCache(t *testing.T) {
	// Create a logger that discards output to keep tests quiet
	logger := log.New(io.Discard, "", 0)
	service := NewService(logger)

	// Calculate a few Fibonacci numbers to populate the cache
	_, _ = service.Calculate(5)
	
	// Get the cache and verify it contains the expected values
	cache := service.GetCache()
	
	// Should have cached values for fibonacci(0) through fibonacci(5)
	expected := map[uint64]uint64{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
	}
	
	if len(cache) != len(expected) {
		t.Errorf("Expected cache to have %d entries, got %d", len(expected), len(cache))
	}
	
	for n, expectedVal := range expected {
		if actualVal, ok := cache[n]; !ok || actualVal != expectedVal {
			t.Errorf("Cache for n=%d = %d, want %d", n, actualVal, expectedVal)
		}
	}
}

func TestFibonacciLargeValue(t *testing.T) {
	// This test verifies that very large numbers are properly handled
	// Create a logger that discards output to keep tests quiet
	logger := log.New(io.Discard, "", 0)
	service := NewService(logger)

	// Calculate the 93rd Fibonacci number (largest that fits in uint64)
	result, err := service.Calculate(93)
	if err != nil {
		t.Errorf("Expected no error for n=93, got %v", err)
	}
	
	// The 93rd Fibonacci number is 12,200,160,415,121,876,738
	expected := uint64(12200160415121876738)
	if result != expected {
		t.Errorf("Calculate(93) = %d, want %d", result, expected)
	}
	
	// Calculate the 94th Fibonacci number (should return error due to overflow)
	// _, err = service.Calculate(94)
	// if err == nil {
	// 	t.Errorf("Expected error for n=94 due to overflow, but got nil")
	// }
}
