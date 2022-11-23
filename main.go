package main

import (
	"fmt"
	"patterns/adapter"
	"patterns/command"
	"patterns/decorator"
	"patterns/dependControl"
	"patterns/facade"
	"patterns/factory"
	"patterns/proxy"
	"patterns/simple_factory"
	"patterns/singleton/atomicOnce"
	"patterns/staticFactory/computer"
	"patterns/staticFactory/fruit"
	"patterns/template"
	"patterns/union"
)

func main() {

	// 依赖反转
	dependControl.BizDriveCar()

	// 练习：组合优于继承
	union.BizUnion()

	// 简单工厂
	simpleFactory.Biz()

	// 练习：工厂模式
	factory.BizFactory()

	// 练习：抽象工厂--生产水果
	fruit.Biz()

	// 练习：抽象工厂--生产电脑配件
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

	// 门面模式
	facade.TestFacade()
	facade.TestKTVFacade()

	// 模板方法模式
	template.TestMakeCoffee()
	template.TestMakeTea()

	// 命令模式
	command.TestCommand()
	command.TestShaoKao()
}
