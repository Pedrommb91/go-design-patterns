package decorator

import (
	"testing"
)

func TestSimpleCoffee(t *testing.T) {
	coffee := &SimpleCoffee{}
	if coffee.Cost() != 2.00 {
		t.Errorf("Expected cost of SimpleCoffee to be 2.00, got %f", coffee.Cost())
	}
	if coffee.Description() != "Simple Coffee" {
		t.Errorf("Expected description of SimpleCoffee to be 'Simple Coffee', got %s", coffee.Description())
	}
}

func TestMilkDecorator(t *testing.T) {
	coffee := &SimpleCoffee{}
	milkCoffee := &MilkDecorator{}
	milkCoffee.Decorate(coffee)
	if milkCoffee.Cost() != 2.50 {
		t.Errorf("Expected cost of MilkDecorator to be 2.50, got %f", milkCoffee.Cost())
	}
	if milkCoffee.Description() != "Simple Coffee, Milk" {
		t.Errorf("Expected description of MilkDecorator to be 'Simple Coffee, Milk', got %s", milkCoffee.Description())
	}
}

func TestSugarDecorator(t *testing.T) {
	coffee := &SimpleCoffee{}
	sugarCoffee := &SugarDecorator{}
	sugarCoffee.Decorate(coffee)
	if sugarCoffee.Cost() != 2.25 {
		t.Errorf("Expected cost of SugarDecorator to be 2.25, got %f", sugarCoffee.Cost())
	}
	if sugarCoffee.Description() != "Simple Coffee, Sugar" {
		t.Errorf("Expected description of SugarDecorator to be 'Simple Coffee, Sugar', got %s", sugarCoffee.Description())
	}
}

func TestCoffeeWorkflow(t *testing.T) {
	// Create a simple coffee
	coffee := &SimpleCoffee{}

	// Decorate it with Milk
	milkCoffee := &MilkDecorator{}
	milkCoffee.Decorate(coffee)
	if milkCoffee.Cost() != 2.50 {
		t.Errorf("Expected cost of MilkDecorator to be 2.50, got %f", milkCoffee.Cost())
	}
	if milkCoffee.Description() != "Simple Coffee, Milk" {
		t.Errorf("Expected description of MilkDecorator to be 'Simple Coffee, Milk', got %s", milkCoffee.Description())
	}

	// Decorate it with Sugar
	sugarCoffee := &SugarDecorator{}
	sugarCoffee.Decorate(milkCoffee)
	if sugarCoffee.Cost() != 2.75 {
		t.Errorf("Expected cost of SugarDecorator to be 2.75, got %f", sugarCoffee.Cost())
	}
	if sugarCoffee.Description() != "Simple Coffee, Milk, Sugar" {
		t.Errorf("Expected description of SugarDecorator to be 'Simple Coffee, Milk, Sugar', got %s", sugarCoffee.Description())
	}
}
