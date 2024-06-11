package builder

import (
	"testing"
)

func TestBuilder_SetName(t *testing.T) {
	product := NewSimpleBuilder().SetName("TestName").Build()

	if product.GetName() != "TestName" {
		t.Errorf("Expected name to be 'TestName', but got %s", product.GetName())
	}
}

func TestBuilder_SetValue(t *testing.T) {
	product := NewSimpleBuilder().SetValue(100).Build()

	if product.GetValue() != 100 {
		t.Errorf("Expected value to be 100, but got %d", product.GetValue())
	}
}

func TestBuilder_Build(t *testing.T) {
	product := NewSimpleBuilder().
		SetName("TestName").
		SetValue(100).
		Build()

	if product.GetName() != "TestName" {
		t.Errorf("Expected name to be 'TestName', but got %s", product.GetName())
	}

	if product.GetValue() != 100 {
		t.Errorf("Expected value to be 100, but got %d", product.GetValue())
	}
}
