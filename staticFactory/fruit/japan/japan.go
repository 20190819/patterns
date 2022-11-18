package japan

import (
	"fmt"
	abstract2 "patterns/staticFactory/fruit/abstract"
)

type AppleJapan struct {
}

func (a *AppleJapan) ShowApple() {
	fmt.Println("日本苹果")
}

type PearJapan struct {
}

func (p *PearJapan) ShowPear() {
	fmt.Println("日本梨子")
}

type BananaJapan struct {
}

func (b *BananaJapan) ShowBanana() {
	fmt.Println("日本香蕉")
}

type FactoryJapan struct {
}

func (f *FactoryJapan) CreateApple() abstract2.AbstractApple {
	return new(AppleJapan)
}

func (f *FactoryJapan) CreatePear() abstract2.AbstractPear {
	return new(PearJapan)
}

func (f *FactoryJapan) CreateBanana() abstract2.AbstractBanana {
	return new(BananaJapan)
}
