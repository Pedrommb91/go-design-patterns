package proxy

// Door interface
type Door interface {
	Open() string
	Close() string
}

// LabDoor struct that implements Door interface
type LabDoor struct{}

// Open method for LabDoor
func (d *LabDoor) Open() string {
	return "Lab door opened"
}

// Close method for LabDoor
func (d *LabDoor) Close() string {
	return "Lab door closed"
}

// Security struct that acts as a proxy
type Security struct {
	door     Door
	password string
}

// NewSecurity creates a new Security
func NewSecurity(door Door, password string) *Security {
	return &Security{
		door:     door,
		password: password,
	}
}

// Authenticate checks if the input password is correct
func (s *Security) Authenticate(inputPassword string) bool {
	return inputPassword == s.password
}

// Open method for Security that checks authentication before opening the door
func (s *Security) Open(inputPassword string) string {
	if s.Authenticate(inputPassword) {
		return s.door.Open()
	}
	return "Access denied"
}

// Close method for Security that delegates the closing of the door
func (s *Security) Close() string {
	return s.door.Close()
}
