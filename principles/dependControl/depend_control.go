package dependControl

import "fmt"

// 依赖反转
// 抽象层 实现层 业务逻辑层

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("bmw run")
}

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("benz run")
}

type Zhang3 struct {
}

func (z3 *Zhang3) Drive(car Car) {
	fmt.Println("zhang3 start drive")
	car.Run()
}

type Li4 struct {
}

func (l4 *Li4) Drive(car Car) {
	fmt.Println("Li4 start drive")
	car.Run()
}

func BizDriveCar() {
	z3 := Zhang3{}
	l4 := Li4{}

	z3.Drive(new(BMW))
	z3.Drive(new(Benz))

	l4.Drive(new(BMW))
	l4.Drive(new(Benz))
}
