package principles

import "fmt"

type ClothesShop struct {
}

func (c *ClothesShop) OnShop() {
	fmt.Println("休闲的装扮")
}

type ClothesWork struct {
}

func (w *ClothesWork) OnWork() {
	fmt.Println("工作服装扮")
}

func TestSingle() {
	shop := new(ClothesShop)
	shop.OnShop()

	work := new(ClothesWork)
	work.OnWork()
}
