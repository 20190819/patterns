package factory

import "fmt"

// 工厂模式

// 抽象层

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

// 实现层

type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("这是苹果")
}

type AppleJapan struct {
}

func (aj *AppleJapan) Show() {
	fmt.Println("这是日本苹果")
}

type Pear struct {
}

func (p *Pear) Show() {
	fmt.Println("这是梨子")
}

type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("这是日本梨子")
}

type AppleFactory struct {
}

func (a *AppleFactory) CreateFruit() Fruit {
	return new(Apple)
}

type PearFactory struct {
}

func (p *PearFactory) CreateFruit() Fruit {
	return new(Pear)
}

type BananaFactory struct {
}

func (f *BananaFactory) CreateFruit() Fruit {
	return new(Banana)
}

type AppleJapanFactory struct {
}

func (aj *AppleJapanFactory) CreateFruit() Fruit {
	return new(AppleJapan)
}

// 业务逻辑层

func BizFactory() {

	fmt.Println("===工厂模式===")

	appleFac := new(AppleFactory)
	apple := appleFac.CreateFruit()
	apple.Show()

	pearFac := new(PearFactory)
	pear := pearFac.CreateFruit()
	pear.Show()

	bananaFac := new(BananaFactory)
	banana := bananaFac.CreateFruit()
	banana.Show()

	// (+)
	appleJapanFac:=new(AppleJapanFactory)
	appleJapan:=appleJapanFac.CreateFruit()
	appleJapan.Show()
}
