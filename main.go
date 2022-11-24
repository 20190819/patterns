package main

import (
	"fmt"
	"patterns/adapter"
	"patterns/command"
	"patterns/decorator"
	"patterns/facade"
	"patterns/factory"
	"patterns/observer"
	dependControl2 "patterns/principles/dependControl"
	union2 "patterns/principles/union"
	"patterns/proxy"
	"patterns/simple_factory"
	"patterns/singleton/atomicOnce"
	"patterns/staticFactory/computer"
	"patterns/staticFactory/fruit"
	"patterns/strategy"
	"patterns/template"
)

func main() {

	// ========设计原则========
	// 单一职责原则

	// 开闭原则

	// 里氏代换原则

	// 接口隔离原则

	// 迪米特法则

	// 依赖反转原则
	dependControl2.BizDriveCar()

	// 原则：合成复用优于继承
	union2.BizUnion()

	//==========常用设计模式==========

	// 简单工厂
	simpleFactory.Biz()

	// 工厂模式
	factory.BizFactory()

	// 抽象工厂--生产水果
	fruit.Biz()

	// 抽象工厂--生产电脑配件
	computer.Biz()

	// 单例模式
	// atomicOnce 标记内存状态
	s1 := atomicOnce.GetInstanceByOnce()
	s2 := atomicOnce.GetInstanceByOnce()
	fmt.Println(s1 == s2)

	// 代理模式
	proxy.TestProxy()

	// 装饰器模式
	decorator.TestDecorator()

	// 适配器模式
	adapter.TestAdapter()

	// 门面(外观)模式
	facade.TestFacade()
	facade.TestKTVFacade()

	// 模板方法模式
	template.TestMakeCoffee()
	template.TestMakeTea()

	// 命令模式
	command.TestCommand()
	command.TestShaoKao()

	// 策略模式
	strategy.TestStrategy()
	strategy.TestSaleStrategy()

	// 观察者模式
	observer.TestStudentObserver()
}
