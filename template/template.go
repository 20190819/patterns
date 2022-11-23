package template

import "fmt"

type Beverage interface {
	BoilWater()          // 煮开水
	Brew()               // 冲泡
	PourInCup()          //倒入杯中
	AddThings()          //添加酌料
	WantAddThings() bool // 是否添加酌料（hook）
}

// Template 流程模板，让具体的制作过程继承且实现
type Template struct {
	b Beverage
}

func (t *Template) MakeBeverage() {
	fmt.Printf("\n")
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	//子类可以覆写该方法，决定是否执行下面的逻辑
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}
