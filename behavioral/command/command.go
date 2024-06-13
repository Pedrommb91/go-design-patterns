package command

import "fmt"

// Command interface declares a method for executing a command
type Command interface {
	Execute()
}

// Light represents a Receiver in the Command pattern
type Light struct {
	IsOn bool
}

// TurnOnCommand is a ConcreteCommand to turn on the light
type TurnOnCommand struct {
	Light *Light
}

// Execute method of TurnOnCommand
func (c *TurnOnCommand) Execute() {
	c.Light.IsOn = true
	fmt.Println("Light turned on")
}

// TurnOffCommand is a ConcreteCommand to turn off the light
type TurnOffCommand struct {
	Light *Light
}

// Execute method of TurnOffCommand
func (c *TurnOffCommand) Execute() {
	c.Light.IsOn = false
	fmt.Println("Light turned off")
}

// Switch represents an Invoker in the Command pattern
type Switch struct {
	commands map[string]Command
}

// NewSwitch creates a new Switch
func NewSwitch() *Switch {
	return &Switch{commands: make(map[string]Command)}
}

// StoreCommand stores a command to be invoked later
func (s *Switch) StoreCommand(key string, command Command) {
	s.commands[key] = command
}

// ExecuteCommand invokes the command associated with a key
func (s *Switch) ExecuteCommand(key string) {
	if command, exists := s.commands[key]; exists {
		command.Execute()
	} else {
		fmt.Printf("Command not found: %s\n", key)
	}
}
