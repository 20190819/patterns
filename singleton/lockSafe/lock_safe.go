package lockSafe

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例之线程安全模式")
}
