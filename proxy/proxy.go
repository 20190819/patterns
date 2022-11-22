package proxy

import "fmt"

type Goods struct {
	Kind string // 种类
	Fact bool   // 真伪
	Safe bool   // 安全
}

// 抽象层

type Shopping interface {
	Buy(goods *Goods)
}

type Proxy interface {
	Buy(goods *Goods)              // 购买
	Distinguish(goods *Goods) bool // 辨别
	Check(goods *Goods) bool       // 安检
}

func prt(country string, goods *Goods) {
	fmt.Printf("去%s购物，购买%v\n", country, goods.Kind)
}

// 实现层

type KoraShopping struct{}

func (k *KoraShopping) Buy(goods *Goods) {
	prt("韩国", goods)
}

type AmericanShopping struct{}

func (a *AmericanShopping) Buy(goods *Goods) {
	prt("美国", goods)
}

type JapanShopping struct{}

func (j *JapanShopping) Buy(goods *Goods) {
	prt("日本", goods)
}

// 业务逻辑层

type OverseasProxy struct {
	shopping Shopping
}

func NewOverseasProxy(shopping Shopping) *OverseasProxy {
	return &OverseasProxy{shopping: shopping}
}

func (o *OverseasProxy) Buy(goods *Goods) {

	// 鉴别
	if !o.Distinguish(goods) {
		fmt.Printf("%s 是假货\n", goods.Kind)
		return
	}

	// 安检
	if !o.Check(goods) {
		fmt.Printf("%s 安检未通过\n", goods.Kind)
		return
	}

	o.shopping.Buy(goods)
}

func (o *OverseasProxy) Distinguish(goods *Goods) bool {
	fmt.Println("进行辨别真伪")
	return goods.Fact
}

func (o *OverseasProxy) Check(goods *Goods) bool {
	fmt.Println("进行安检")
	return goods.Safe
}

func TestProxy() {
	goods1 := &Goods{
		Kind: "面膜",
		Fact: true,
		Safe: true,
	}
	goods2 := &Goods{
		Kind: "西瓜",
		Fact: false,
		Safe: false,
	}

	proxy1 := NewOverseasProxy(new(KoraShopping))

	proxy1.Buy(goods1)
	proxy1.Buy(goods2)
}
