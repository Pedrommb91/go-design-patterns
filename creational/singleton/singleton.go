// singleton.go

// Package singleton demonstrates the Singleton Design Pattern in Go.
// The Singleton pattern ensures that a class has only one instance and provides a global point of access to it.
// This implementation uses sync.Once to ensure that the instance is created only once in a thread-safe manner.

package singleton

import "sync"

// Singleton is the type for singleton instance
type Singleton struct {
	// Properties of the singleton
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
