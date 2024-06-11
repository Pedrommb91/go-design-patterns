package builder

import (
	"testing"
)

func TestDefaultBuilder(t *testing.T) {
	product := NewBuilder(Default).Build()

	if product.name != "Default" {
		t.Errorf("Expected product name to be 'Default', got '%s'", product.name)
	}

	if product.value != 0 {
		t.Errorf("Expected product value to be 0, got %d", product.value)
	}
}

func TestCustomBuilder(t *testing.T) {
	builder := NewBuilder(Custom)
	product := builder.SetName("CustomProduct").SetValue(20).Build()

	if product.name != "CustomProduct" {
		t.Errorf("Expected product name to be 'CustomProduct', got '%s'", product.name)
	}

	if product.value != 20 {
		t.Errorf("Expected product value to be 20, got %d", product.value)
	}
}

func TestNewBuilder(t *testing.T) {
	defaultBuilder := NewBuilder(Default)
	if defaultBuilder == nil {
		t.Error("Expected defaultBuilder to be non-nil")
	}

	customBuilder := NewBuilder(Custom)
	if customBuilder == nil {
		t.Error("Expected customBuilder to be non-nil")
	}

	invalidBuilder := NewBuilder(BuilderType(999))
	if invalidBuilder != nil {
		t.Error("Expected invalidBuilder to be nil")
	}
}
