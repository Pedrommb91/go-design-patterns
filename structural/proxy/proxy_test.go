package proxy

import (
	"testing"
)

func TestLabDoor(t *testing.T) {
	door := &LabDoor{}

	openMsg := door.Open()
	if openMsg != "Lab door opened" {
		t.Errorf("Expected 'Lab door opened', got '%s'", openMsg)
	}

	closeMsg := door.Close()
	if closeMsg != "Lab door closed" {
		t.Errorf("Expected 'Lab door closed', got '%s'", closeMsg)
	}
}

func TestSecurity(t *testing.T) {
	door := &LabDoor{}
	security := NewSecurity(door, "abc123")

	// Test authentication failure
	authFailMsg := security.Open("wrong")
	if authFailMsg != "Access denied" {
		t.Errorf("Expected 'Access denied', got '%s'", authFailMsg)
	}

	// Test authentication success
	authSuccessMsg := security.Open("abc123")
	if authSuccessMsg != "Lab door opened" {
		t.Errorf("Expected 'Lab door opened', got '%s'", authSuccessMsg)
	}

	// Test close
	closeMsg := security.Close()
	if closeMsg != "Lab door closed" {
		t.Errorf("Expected 'Lab door closed', got '%s'", closeMsg)
	}
}
