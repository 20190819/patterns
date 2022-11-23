package template

import "fmt"

type Tea struct {
}

func (t *Tea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

func (t *Tea) Brew() {
	fmt.Println("用水冲茶叶")
}

func (t *Tea) PourInCup() {
	fmt.Println("将充好的咖啡倒入茶壶中")
}

func (t *Tea) AddThings() {
	//fmt.Println("添加柠檬")
}

func (t *Tea) WantAddThings() bool {
	return false
}

type MakeTea struct {
	Template
}

func NewMakeTea(b Beverage) *MakeTea {
	return &MakeTea{Template{b}}
}

func TestMakeTea() {
	mt := NewMakeTea(new(Tea))
	mt.MakeBeverage()
}
