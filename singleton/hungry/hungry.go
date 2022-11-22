package hungry

import "fmt"

type singleton struct{}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例-饿汉模式")
}
