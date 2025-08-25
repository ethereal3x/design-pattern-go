package singleton

import (
	"fmt"
	"sync"
)

var (
	lazy *Lazy
	once sync.Once
)

type Lazy struct {
}

func (lz *Lazy) Hello() {
	fmt.Print("Hi!")
}

func GetLazy() *Lazy {
	once.Do(func() {
		lazy = &Lazy{}
	})
	return lazy
}
