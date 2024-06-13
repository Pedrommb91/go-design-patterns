package chain

import "fmt"

// Approver is the Handler interface in the Chain of Responsibility pattern
type Approver interface {
	ApproveRequest(amount float64) (approved bool)
	SetNext(approver Approver) Approver
}

// BaseApprover provides a default implementation of the Approver interface
type BaseApprover struct {
	next Approver
}

// SetNext sets the next approver in the chain
func (a *BaseApprover) SetNext(approver Approver) Approver {
	a.next = approver
	return approver
}

// ApproveRequest in BaseApprover delegates the request to the next approver
func (a *BaseApprover) ApproveRequest(amount float64) (approved bool) {
	if a.next != nil {
		return a.next.ApproveRequest(amount)
	}
	return false // No approver available in the chain
}

// Manager handles small purchase requests
type Manager struct {
	BaseApprover
}

// ApproveRequest method for Manager
func (m *Manager) ApproveRequest(amount float64) (approved bool) {
	if amount <= 1000 {
		// Manager can approve requests up to $1000
		fmt.Println("Manager approved request.")
		return true
	}
	// Delegate to the next approver
	return m.BaseApprover.ApproveRequest(amount)
}

// Director handles medium purchase requests
type Director struct {
	BaseApprover
}

// ApproveRequest method for Director
func (d *Director) ApproveRequest(amount float64) (approved bool) {
	if amount <= 5000 {
		// Director can approve requests up to $5000
		fmt.Println("Director approved request.")
		return true
	}
	// Delegate to the next approver
	return d.BaseApprover.ApproveRequest(amount)
}

// VicePresident handles large purchase requests
type VicePresident struct {
	BaseApprover
}

// ApproveRequest method for VicePresident
func (vp *VicePresident) ApproveRequest(amount float64) (approved bool) {
	if amount <= 10000 {
		// VicePresident can approve requests up to $10000
		fmt.Println("Vice President approved request.")
		return true
	}
	// Delegate to the next approver
	return vp.BaseApprover.ApproveRequest(amount)
}

// CEO handles all other purchase requests
type CEO struct {
	BaseApprover
}

// ApproveRequest method for CEO
func (c *CEO) ApproveRequest(amount float64) (approved bool) {
	// CEO can approve all requests
	fmt.Println("CEO approved request.")
	return true
}
