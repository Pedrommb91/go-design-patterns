package factory

// Chair is the interface for chair products
type Chair interface {
	SitOn() string
}

// Table is the interface for table products
type Table interface {
	PlaceItems() string
}

// ModernChair is a concrete implementation of Chair
type ModernChair struct{}

func (c *ModernChair) SitOn() string {
	return "Sitting on a Modern Chair"
}

// VictorianChair is a concrete implementation of Chair
type VictorianChair struct{}

func (c *VictorianChair) SitOn() string {
	return "Sitting on a Victorian Chair"
}

// ModernTable is a concrete implementation of Table
type ModernTable struct{}

func (t *ModernTable) PlaceItems() string {
	return "Placing items on a Modern Table"
}

// VictorianTable is a concrete implementation of Table
type VictorianTable struct{}

func (t *VictorianTable) PlaceItems() string {
	return "Placing items on a Victorian Table"
}

// FurnitureFactory is the interface for the abstract factory
type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
}

func GetFurnitureFactory(style string) FurnitureFactory {
	if style == "modern" {
		return &ModernFurnitureFactory{}
	} else if style == "victorian" {
		return &VictorianFurnitureFactory{}
	}
	return nil
}

// ModernFurnitureFactory is a concrete implementation of FurnitureFactory
type ModernFurnitureFactory struct{}

func (f *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (f *ModernFurnitureFactory) CreateTable() Table {
	return &ModernTable{}
}

// VictorianFurnitureFactory is a concrete implementation of FurnitureFactory
type VictorianFurnitureFactory struct{}

func (f *VictorianFurnitureFactory) CreateChair() Chair {
	return &VictorianChair{}
}

func (f *VictorianFurnitureFactory) CreateTable() Table {
	return &VictorianTable{}
}
