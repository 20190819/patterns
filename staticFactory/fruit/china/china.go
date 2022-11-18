package china

import (
	"fmt"
	abstract2 "patterns/staticFactory/fruit/abstract"
)

type AppleChina struct {
}

func (c *AppleChina) ShowApple() {
	fmt.Println("中国苹果")
}

type PearChina struct {
}

func (p *PearChina) ShowPear() {
	fmt.Println("中国梨子")
}

type BananaChina struct {
}

func (b *BananaChina) ShowBanana() {
	fmt.Println("中国香蕉")
}

type FactoryChina struct {
}

func (f *FactoryChina) CreateApple() abstract2.AbstractApple {
	return new(AppleChina)
}

func (f *FactoryChina) CreatePear() abstract2.AbstractPear {
	return new(PearChina)
}
func (f *FactoryChina) CreateBanana() abstract2.AbstractBanana {
	return new(BananaChina)
}
