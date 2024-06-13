package chain

import (
	"testing"
)

// TestChainOfResponsibility tests the chain of approvers
func TestChainOfResponsibility(t *testing.T) {
	// Set up the chain of approvers
	manager := &Manager{}
	director := &Director{}
	vicePresident := &VicePresident{}
	ceo := &CEO{}

	manager.SetNext(director).SetNext(vicePresident).SetNext(ceo)

	// Test cases
	var tests = []struct {
		amount   float64
		approved bool
	}{
		{500, true},   // Manager can approve
		{1500, true},  // Director can approve
		{5500, true},  // VicePresident can approve
		{10500, true}, // CEO can approve
		{11000, true}, // CEO can approve any amount
	}

	for _, tt := range tests {
		approved := manager.ApproveRequest(tt.amount)
		if approved != tt.approved {
			t.Errorf("ApproveRequest(%f): expected %v, actual %v", tt.amount, tt.approved, approved)
		}
	}
}
