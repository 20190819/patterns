package nvidia

import (
	"fmt"
	"patterns/staticFactory/computer/abstract"
)

type ComputerFactoryNvidia struct {
}

func (c *ComputerFactoryNvidia) CreateStorage() abstract.Storage {
	return new(StorageNvidia)
}

func (c *ComputerFactoryNvidia) CreateDisplay() abstract.Display {
	return new(DisplayNvidia)
}

func (c *ComputerFactoryNvidia) CreateCalculate() abstract.Calculate {
	return new(CalculateNvidia)
}

type StorageNvidia struct {
}

func (s *StorageNvidia) ShowStorage() {
	fmt.Println("NVIDIA 内存")
}

type DisplayNvidia struct {
}

func (d *DisplayNvidia) ShowDisplay() {
	fmt.Println("NVIDIA 显卡")
}

type CalculateNvidia struct {
}

func (c *CalculateNvidia) ShowCalculate() {
	fmt.Println("NVIDIA CPU")
}
