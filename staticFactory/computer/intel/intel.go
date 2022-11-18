package intel

import (
	"fmt"
	"patterns/staticFactory/computer/abstract"
)

type ComputerFactoryIntel struct {
}

func (c *ComputerFactoryIntel) CreateStorage() abstract.Storage {
	return new(StorageIntel)
}

func (c *ComputerFactoryIntel) CreateDisplay() abstract.Display {
	return new(DisplayIntel)
}
func (c *ComputerFactoryIntel) CreateCalculate() abstract.Calculate {
	return new(CalculateIntel)
}

type StorageIntel struct {
}

func (s *StorageIntel) ShowStorage() {
	fmt.Println("因特尔内存")
}

type DisplayIntel struct {
}

func (d *DisplayIntel) ShowDisplay() {
	fmt.Println("因特尔显卡")
}

type CalculateIntel struct {
}

func (c *CalculateIntel) ShowCalculate() {
	fmt.Println("因特尔CPU")
}
