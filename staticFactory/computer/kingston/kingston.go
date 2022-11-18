package kingston

import (
	"fmt"
	"patterns/staticFactory/computer/abstract"
)

type ComputerFactoryKingston struct {
}

func (c *ComputerFactoryKingston) CreateStorage() abstract.Storage {
	return new(StorageKingston)
}

func (c *ComputerFactoryKingston) CreateDisplay() abstract.Display {
	return new(DisplayKingston)
}

func (c *ComputerFactoryKingston) CreateCalculate() abstract.Calculate {
	return new(CalculateKingston)
}

type StorageKingston struct {
}

func (s *StorageKingston) ShowStorage() {
	fmt.Println("金士顿内存")
}

type DisplayKingston struct {
}

func (i *DisplayKingston) ShowDisplay() {
	fmt.Println("金士顿显卡")
}

type CalculateKingston struct {
}

func (c *CalculateKingston) ShowCalculate() {
	fmt.Println("金士顿CPU")
}
