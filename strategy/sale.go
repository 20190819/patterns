package strategy

import (
	"fmt"
)

//练习：
//商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景

// SaleStrategy 促销抽象策略
type SaleStrategy interface {
	Discount(price float64) float64
}

// AStrategy 促销具体策略
type AStrategy struct {
}

// Discount 原价*八折
func (a *AStrategy) Discount(price float64) float64 {
	return price * 0.8
}

// BStrategy 促销具体策略
type BStrategy struct {
}

// Discount 满 200 减 100
func (b *BStrategy) Discount(price float64) float64 {
	if price >= 200 {
		return price - 100
	}
	return price
}

// 环境类（使用算法的角色，它在解决问题时，可以采用不同策略）

type Goods struct {
	stg   SaleStrategy
	price float64
}

func (g *Goods) SetStrategy(strategy SaleStrategy) {
	g.stg = strategy
}

// SalePrice 商品售价
func (g *Goods) SalePrice() {
	fmt.Printf("促销后的价格：%.2f\n", g.stg.Discount(g.price))
}

func TestSaleStrategy() {
	goods1 := &Goods{price: 135}
	goods1.SetStrategy(new(AStrategy))
	goods1.SalePrice()

	goods2 := &Goods{price: 202}
	goods2.SetStrategy(new(BStrategy))
	goods2.SalePrice()
}
