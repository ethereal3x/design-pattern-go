package bridge

import "fmt"

type Applications interface {
	Run()
	Name() string
}

type Bulb struct {
}

func (b *Bulb) Run() {
	fmt.Println("Bulb is running")
}

func (b *Bulb) Name() string {
	return "Bulb"
}

type Fan struct {
}

func (f *Fan) Name() string {
	return "Fan"
}

func (f *Fan) Run() {
	fmt.Printf("电扇 %v 转动\n", f.Name())
}

type Switch interface {
	TurnOn()
	TurnOff()
}

type CircleSwitch struct {
	application Applications
}

func NewCircleSwitch(application Applications) *CircleSwitch {
	return &CircleSwitch{application}
}

func (c *CircleSwitch) TurnOn() {
	fmt.Printf("打开圆形开关，打开 %s，开始运行 ", c.application.Name())
	c.application.Run()
}

func (c *CircleSwitch) TurnOff() {
	fmt.Printf("关闭圆形开关，关闭 %s\n", c.application.Name())
}

type SquareSwitch struct {
	application Applications // 关联的家电
}

func NewSquareSwitch(application Applications) *SquareSwitch {
	return &SquareSwitch{application: application}
}

func (c *SquareSwitch) TurnOn() {
	fmt.Printf("打开方形开关，打开 %s，开始运行 ", c.application.Name())
	c.application.Run()
}

func (c *SquareSwitch) TurnOff() {
	fmt.Printf("关闭方形开关，关闭 %s\n", c.application.Name())
}
