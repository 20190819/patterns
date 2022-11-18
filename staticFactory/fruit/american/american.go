package american

import (
	"fmt"
	abstract2 "patterns/staticFactory/fruit/abstract"
)

type AppleAmerican struct {
}

func (a *AppleAmerican) ShowApple() {
	fmt.Println("美国苹果")
}

type PearAmerican struct {
}

func (p *PearAmerican) ShowPear() {
	fmt.Println("美国梨子")
}

type BananaAmerican struct {

}

func (b *BananaAmerican) ShowBanana()  {
	fmt.Println("美国香蕉")
}

type FactoryAmerican struct {
}

func (f *FactoryAmerican) CreateApple() abstract2.AbstractApple {
	return new(AppleAmerican)
}

func (f *FactoryAmerican) CreatePear() abstract2.AbstractPear {
	return new(PearAmerican)
}

func (f *FactoryAmerican) CreateBanana() abstract2.AbstractBanana {
	return new(BananaAmerican)
}