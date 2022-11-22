package atomicOnce

import (
	"fmt"
	"sync"
)

var once sync.Once

type singletonByOnce struct{}

var instanceOnce *singletonByOnce

func GetInstanceByOnce() *singletonByOnce {
	once.Do(func() {
		instanceOnce = new(singletonByOnce)
	})
	return instanceOnce
}

func (s *singletonByOnce) SomeThing() {
	fmt.Println("利用 sync/once 实现单例")
}
