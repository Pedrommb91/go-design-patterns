package factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	tests := []struct {
		factoryType string
		expectedA   string
		expectedB   string
	}{
		{"1", "Using Product A1", "Using Product B1"},
		{"2", "Using Product A2", "Using Product B2"},
	}

	for _, test := range tests {
		factory := GetAbstractFactory(test.factoryType)
		if factory == nil {
			t.Fatalf("expected a valid factory for type %v, but got nil", test.factoryType)
		}

		productA := factory.CreateProductA()
		if productA.UseA() != test.expectedA {
			t.Errorf("expected productA use to be %v, but got %v", test.expectedA, productA.UseA())
		}

		productB := factory.CreateProductB()
		if productB.UseB() != test.expectedB {
			t.Errorf("expected productB use to be %v, but got %v", test.expectedB, productB.UseB())
		}
	}
}
