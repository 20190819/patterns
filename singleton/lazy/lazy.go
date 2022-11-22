package lazy

import "fmt"

type lazy struct{}

var instance *lazy

func GetInstance() *lazy {
	if instance == nil {
		instance = new(lazy)
	}
	return instance
}

func (l *lazy) SomeThing() {
	fmt.Println("单例之懒汉模式")
}
