package template

import "fmt"

type Coffee struct {
}

func (c *Coffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (c *Coffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}

func (c *Coffee) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

func (c *Coffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (c *Coffee) WantAddThings() bool {
	return true
}

type MakeCoffee struct {
	Template
}

func NewMakeCoffee(b Beverage) *MakeCoffee {
	return &MakeCoffee{Template{b}}
}

func TestMakeCoffee() {
	mc := NewMakeCoffee(new(Coffee))
	mc.MakeBeverage()
}
