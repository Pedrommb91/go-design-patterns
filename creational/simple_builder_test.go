package creational

import (
	"testing"
)

func TestBuilder_SetName(t *testing.T) {
	builder := NewBuilder()
	builder.SetName("TestName")
	product := builder.Build()

	if product.GetName() != "TestName" {
		t.Errorf("Expected name to be 'TestName', but got %s", product.GetName())
	}
}

func TestBuilder_SetValue(t *testing.T) {
	builder := NewBuilder()
	builder.SetValue(100)
	product := builder.Build()

	if product.GetValue() != 100 {
		t.Errorf("Expected value to be 100, but got %d", product.GetValue())
	}
}

func TestBuilder_Build(t *testing.T) {
	builder := NewBuilder()
	builder.SetName("TestName")
	builder.SetValue(100)
	product := builder.Build()

	if product.GetName() != "TestName" {
		t.Errorf("Expected name to be 'TestName', but got %s", product.GetName())
	}

	if product.GetValue() != 100 {
		t.Errorf("Expected value to be 100, but got %d", product.GetValue())
	}
}
