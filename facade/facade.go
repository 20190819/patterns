package facade

import "fmt"

type SubsystemA struct {
}

func (s *SubsystemA) mA() {
	fmt.Println("this is ma")
}

type SubSystemB struct {
}

func (s *SubSystemB) mB() {
	fmt.Println("this is mb")
}

type SubSystemC struct {
}

func (s *SubSystemC) mC() {
	fmt.Println("this is mc")
}

type SubSystemD struct {
}

func (s *SubSystemD) mD() {
	fmt.Println("this is md")
}

type SystemFacade struct {
	a *SubsystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func NewSystemFacade(a *SubsystemA, b *SubSystemB, c *SubSystemC, d *SubSystemD) *SystemFacade {
	return &SystemFacade{a, b, c, d}
}

func (s *SystemFacade) mAB() {
	s.a.mA()
	s.b.mB()
}

func (s *SystemFacade) mCD() {
	s.c.mC()
	s.d.mD()
}

func TestFacade() {

	fmt.Printf("\n")
	// 不使用 门面模式
	subA := new(SubsystemA)
	subB := new(SubSystemB)

	subA.mA()
	subB.mB()

	// 使用 门面模式
	sysFacade := NewSystemFacade(new(SubsystemA), new(SubSystemB), new(SubSystemC), new(SubSystemD))
	sysFacade.mAB()

}
