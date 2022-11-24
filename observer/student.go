package observer

import (
	"fmt"
)

// 抽象层

type StudentObserver interface {
	DoBadThing()      // 做坏事
	OnTeacherComing() // 老师来了
}

type StudentSubject interface {
	AddObserver(observer StudentObserver)
	RemoveObserver(observer StudentObserver)
	Notify()
}

// 实现层

type Student struct {
	name     string
	badThing string
}

type Zhang3 struct {
	Student
}

func CommonDo(name, bad string) {
	fmt.Printf("%s正在%s\n", name, bad)
}

func CommonStop(name, bad string) {
	fmt.Printf("%s停止%s\n", name, bad)
}

func (z *Zhang3) DoBadThing() {
	CommonDo(z.name, z.badThing)
}

func (z *Zhang3) OnTeacherComing() {
	CommonStop(z.name, z.badThing)
}

type Li4 struct {
	Student
}

func (l *Li4) DoBadThing() {
	CommonDo(l.name, l.badThing)
}

func (l *Li4) OnTeacherComing() {
	CommonDo(l.name, l.badThing)
}

type Wang5 struct {
	Student
}

func (w *Wang5) DoBadThing() {
	CommonDo(w.name, w.badThing)
}

func (w *Wang5) OnTeacherComing() {
	CommonDo(w.name, w.badThing)
}

type ClassMonitor struct {
	observers []StudentObserver
}

func (c *ClassMonitor) AddObserver(observer StudentObserver) {
	c.observers = append(c.observers, observer)
}

func (c *ClassMonitor) RemoveObserver(observer StudentObserver) {
	for i, studentObserver := range c.observers {
		if studentObserver == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
		}
	}
}

func (c *ClassMonitor) Notify() {
	for _, observer := range c.observers {
		observer.OnTeacherComing()
	}
}

func TestStudentObserver() {
	fmt.Printf("\n")
	zhang3 := &Zhang3{Student{"张三", "抄作业"}}
	li4 := &Zhang3{Student{"李四", "打游戏"}}
	wang5 := &Zhang3{Student{"王五", "看别人打游戏"}}
	monitor := new(ClassMonitor)
	monitor.AddObserver(zhang3)
	monitor.AddObserver(li4)
	monitor.AddObserver(wang5)

	fmt.Println("老师不在教室...")
	zhang3.DoBadThing()
	li4.DoBadThing()
	wang5.DoBadThing()

	fmt.Println("老师回来了...")
	monitor.Notify()
}
