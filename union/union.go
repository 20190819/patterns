package union

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("cat can eat2")
}

// 继承的方式
type catB struct {
	Cat
}

func (cb catB) Sleep() {
	fmt.Println("catB can sleep")
}

// 组合的方式
type catC struct {
	cat *Cat
}

func NewCatC() *catC {
	return &catC{cat: new(Cat)}
}

func (cc *catC) Sleep() {
	fmt.Println("catC can sleep")
}

func (cc *catC) Eat() {
	fmt.Println("catC can eat")
}

func BizUnion() {
	// 继承
	cb:=new(catB)
	cb.Eat()
	cb.Sleep()

	// 组合
	cc:=NewCatC()
	cc.cat.Eat()
	cc.Eat()
	cc.Sleep()
}