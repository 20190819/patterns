package adapter

import "fmt"

// 抽象层

// V5 适配目标
type V5 interface {
	Use5V()
}

// 实现层

// V220 适配者
type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("接入 220V 电源")
}

// Adapter 适配器
type Adapter struct {
	v220 *V220
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func (a *Adapter) Use5V() {
	a.v220.Use220V()
	fmt.Println("输出 5V >>>")
}

type Phone struct {
	v5 V5
}

func NewPhone(v5 V5) *Phone {
	return &Phone{v5}
}

func (p *Phone) charge() {
	p.v5.Use5V()
	fmt.Println("phone 充电中 ...")
}

// 业务逻辑层

func TestAdapter() {
	fmt.Println("\n")
	phone := NewPhone(NewAdapter(new(V220)))
	phone.charge()
}
