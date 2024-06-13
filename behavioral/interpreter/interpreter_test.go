package interpreter

import (
	"testing"
)

func TestVariableExpression(t *testing.T) {
	ctx := &Context{
		VariableValues: map[string]int{"x": 5, "y": 10},
	}
	varExp := &VariableExpression{Name: "x"}
	result := varExp.Interpret(ctx)
	if result != 5 {
		t.Errorf("VariableExpression.Interpret() = %d; want 5", result)
	}
}

func TestAddExpression(t *testing.T) {
	ctx := &Context{
		VariableValues: map[string]int{"x": 5, "y": 10},
	}
	addExp := &AddExpression{
		Left:  &VariableExpression{Name: "x"},
		Right: &VariableExpression{Name: "y"},
	}
	result := addExp.Interpret(ctx)
	if result != 15 {
		t.Errorf("AddExpression.Interpret() = %d; want 15", result)
	}
}

func TestSubtractExpression(t *testing.T) {
	ctx := &Context{
		VariableValues: map[string]int{"x": 10, "y": 5},
	}
	subExp := &SubtractExpression{
		Left:  &VariableExpression{Name: "x"},
		Right: &VariableExpression{Name: "y"},
	}
	result := subExp.Interpret(ctx)
	if result != 5 {
		t.Errorf("SubtractExpression.Interpret() = %d; want 5", result)
	}
}
