package command

import "fmt"

//练习：
//联想路边撸串烧烤场景， 有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员MM。根据命令模式，设计烤串场景。

type Shaokao interface {
	Kao()
}

type Cook struct {
}

func (c *Cook) KaoMutton() {
	fmt.Println("烤羊肉")
}

func (c *Cook) KaoChickenWings() {
	fmt.Println("烤鸡翅")
}

type KaoMuttonCommand struct {
	cook *Cook
}

func (kmc *KaoMuttonCommand) Kao() {
	kmc.cook.KaoMutton()
}

type KaoChickenWingsCommand struct {
	cook *Cook
}

func (kcw *KaoChickenWingsCommand) Kao() {
	kcw.cook.KaoChickenWings()
}

type Waiter struct {
	cmds []Shaokao
}

func (w *Waiter) Notify() {
	if len(w.cmds) == 0 {
		return
	}

	for _, cmd := range w.cmds {
		cmd.Kao()
	}
}

func TestShaoKao() {
	cook := new(Cook)
	kmc := &KaoMuttonCommand{cook}
	kcw := &KaoChickenWingsCommand{cook}

	waiter := new(Waiter)
	waiter.cmds = append(waiter.cmds, kmc)
	waiter.cmds = append(waiter.cmds, kcw)
	waiter.Notify()
}
