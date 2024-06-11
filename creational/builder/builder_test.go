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

func TestConcreteBuilder_SetColor(t *testing.T) {
	product := NewBuilder(Concrete).SetColor("Red").Build()

	if product.color != "Red" {
		t.Errorf("Expected color to be 'Red', but got %s", product.color)
	}
}

func TestConcreteBuilder_SetSize(t *testing.T) {
	product := NewBuilder(Concrete).SetSize("Large").Build()

	if product.size != "Large" {
		t.Errorf("Expected size to be 'Large', but got %s", product.size)
	}
}

func TestConcreteBuilder_Build(t *testing.T) {
	product := NewBuilder(Concrete).
		SetName("TestName").
		SetValue(100).
		SetColor("Red").
		SetSize("Large").
		Build()

	if product.name != "TestName" {
		t.Errorf("Expected name to be 'TestName', but got %s", product.name)
	}

	if product.value != 100 {
		t.Errorf("Expected value to be 100, but got %d", product.value)
	}

	if product.color != "Red" {
		t.Errorf("Expected color to be 'Red', but got %s", product.color)
	}

	if product.size != "Large" {
		t.Errorf("Expected size to be 'Large', but got %s", product.size)
	}
}
