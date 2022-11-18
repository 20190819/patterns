package simpleFactory

import "fmt"

type Fruit interface {
	ShowName()
}

type Apple struct {
}

func (a *Apple) ShowName() {
	fmt.Println("这是苹果")
}

type Pear struct {
}

func (p *Pear) ShowName() {
	fmt.Println("这是梨子")
}

type Banana struct {
}

func (b *Banana) ShowName() {
	fmt.Println("这是香蕉")
}

type Factory struct {
}

func (f *Factory) CreateFruit(kind string) Fruit {

	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	}

	if kind == "pear" {
		fruit = new(Pear)
	}

	if kind == "banana" {
		fruit = new(Banana)
	}
	return fruit
}

func Biz() {
	fmt.Println("===简单工厂===")
	fac := new(Factory)
	fac.CreateFruit("apple").ShowName()
	fac.CreateFruit("pear").ShowName()
	fac.CreateFruit("banana").ShowName()
}
