// singleton_test.go

package singleton

import (
	"sync"
	"testing"
)

// TestSingleton tests the singleton instance
func TestSingleton(t *testing.T) {
	// Get the first instance
	instance1 := GetInstance()

	// Get the second instance
	instance2 := GetInstance()

	// Check if both instances are the same
	if instance1 != instance2 {
		t.Errorf("Expected instance1 and instance2 to be the same. Got %p and %p", instance1, instance2)
	}
}

// TestSingletonConcurrent tests the singleton instance in a concurrent scenario
func TestSingletonConcurrent(t *testing.T) {
	// Start a number of goroutines all trying to get the instance
	var wg sync.WaitGroup
	const numGoroutines = 100
	instances := make([]*Singleton, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance()
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check if all instances are the same
	for i := 1; i < numGoroutines; i++ {
		if instances[i] != instances[0] {
			t.Errorf("Expected all instances to be the same. Instance at index 0: %p, instance at index %d: %p", instances[0], i, instances[i])
		}
	}
}
