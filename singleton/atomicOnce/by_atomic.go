package atomicOnce

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var initialized uint32
var lock sync.Mutex

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {

	// 标记位被设置，直接返回
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()
	if initialized == 0 {
		instance = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("使用 atomic 标记内存状态存留")
}
