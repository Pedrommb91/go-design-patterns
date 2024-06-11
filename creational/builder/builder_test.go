package builder

import (
	"testing"
)

func TestConcreteBuilder_SetName(t *testing.T) {
	product := NewBuilder(Concrete).SetName("TestName").Build()

	if product.name != "TestName" {
		t.Errorf("Expected name to be 'TestName', but got %s", product.name)
	}
}

func TestConcreteBuilder_SetValue(t *testing.T) {
	product := NewBuilder(Concrete).SetValue(100).Build()

	if product.value != 100 {
		t.Errorf("Expected value to be 100, but got %d", product.value)
	}
}

func TestConcreteBuilder_Build(t *testing.T) {
	product := NewBuilder(Concrete).
		SetName("TestName").
		SetValue(100).
		Build()

	if product.name != "TestName" {
		t.Errorf("Expected name to be 'TestName', but got %s", product.name)
	}

	if product.value != 100 {
		t.Errorf("Expected value to be 100, but got %d", product.value)
	}
}
