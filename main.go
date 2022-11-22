package main

import (
	"fmt"
	"patterns/dependControl"
	"patterns/factory"
	"patterns/proxy"
	"patterns/simple_factory"
	"patterns/singleton/atomicOnce"
	"patterns/staticFactory/computer"
	"patterns/staticFactory/fruit"
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

	// atomicOnce 标记内存状态
	s1 := atomicOnce.GetInstanceByOnce()
	s2 := atomicOnce.GetInstanceByOnce()
	fmt.Println(s1 == s2)
	s1.SomeThing()
	s2.SomeThing()

	// 代理模式
	proxy.TestProxy()
}
