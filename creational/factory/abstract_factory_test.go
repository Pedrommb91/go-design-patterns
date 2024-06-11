package factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	tests := []struct {
		factoryStyle  string
		expectedChair string
		expectedTable string
	}{
		{"modern", "Sitting on a Modern Chair", "Placing items on a Modern Table"},
		{"victorian", "Sitting on a Victorian Chair", "Placing items on a Victorian Table"},
	}

	for _, test := range tests {
		factory := GetFurnitureFactory(test.factoryStyle)
		if factory == nil {
			t.Fatalf("expected a valid factory for style %v, but got nil", test.factoryStyle)
		}

		chair := factory.CreateChair()
		if chair.SitOn() != test.expectedChair {
			t.Errorf("expected chair sit to be %v, but got %v", test.expectedChair, chair.SitOn())
		}

		table := factory.CreateTable()
		if table.PlaceItems() != test.expectedTable {
			t.Errorf("expected table place items to be %v, but got %v", test.expectedTable, table.PlaceItems())
		}
	}
}
