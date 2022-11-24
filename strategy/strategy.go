package strategy

import "fmt"

// WeaponStrategy 抽象策略类
type WeaponStrategy interface {
	UseWeapon()
}

type AK47 struct {
}

func (a *AK47) UseWeapon() {
	fmt.Println("use AK47")
}

type Knife struct {
}

func (k *Knife) UseWeapon() {
	fmt.Println("use Knife")
}

// 环境类（使用算法的角色，它在解决问题时，可以采用多种策略）

type Hero struct {
	strategy WeaponStrategy
}

func (h *Hero) SetStrategy(strategy WeaponStrategy) {
	h.strategy = strategy
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func TestStrategy() {
	fmt.Printf("\n")
	hero := new(Hero)
	hero.SetStrategy(new(AK47))
	hero.Fight()

	hero.SetStrategy(new(Knife))
	hero.Fight()
}
