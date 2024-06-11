package factory

import (
	"testing"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		productType ProductType
		expected    string
		shouldError bool
	}{
		{ProductA, "Using Product A", false},
		{ProductB, "Using Product B", false},
		{ProductType(999), "", true},
	}

	for _, test := range tests {
		product, err := CreateProduct(test.productType)
		if test.shouldError {
			if err == nil {
				t.Errorf("expected an error for product type %v, but got none", test.productType)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error for product type %v, but got %v", test.productType, err)
			}
			if product.Use() != test.expected {
				t.Errorf("expected product use to be %v, but got %v", test.expected, product.Use())
			}
		}
	}
}
