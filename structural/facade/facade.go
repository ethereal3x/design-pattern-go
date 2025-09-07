package facade

import "fmt"

type VegVendor struct {
}

func (v *VegVendor) Purchase() {
	fmt.Println("蔬菜")
}

type Chef struct{}

func (c *Chef) Cook() {
	fmt.Println("下厨烹饪")
}

type Waiter struct{}

func (h *Waiter) Service() {
	fmt.Println("开始为用户服务")
}

type Cleaner struct{}

func (c *Cleaner) Clean() {
	fmt.Println("开始清洁卫生")
}

type Facade struct {
	vendor  *VegVendor
	chef    *Chef
	waiter  *Waiter
	cleaner *Cleaner
}

func NewFacade() *Facade {
	facade := &Facade{}
	facade.cleaner = &Cleaner{}
	facade.chef = &Chef{}
	facade.waiter = &Waiter{}
	facade.cleaner = &Cleaner{}

	return facade
}

func (f *Facade) Order() {
	// 服务员接待入座
	f.waiter.Service()
	// 厨师开始烹饪
	f.chef.Cook()
	// 上菜
	f.waiter.Service()
	// 收拾桌子洗碗
	f.cleaner.Clean()
}
