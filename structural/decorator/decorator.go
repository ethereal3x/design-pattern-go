package decorator

import "fmt"

type ShowAble interface {
	Show()
}

type Girl struct{}

func NewGirl() *Girl {
	return &Girl{}
}

func (g Girl) Show() {
	fmt.Print("girl simple")
}

type Decorator struct {
	showAble ShowAble
}

func NewDecorator(showAble ShowAble) *Decorator {
	return &Decorator{showAble}
}

func (d *Decorator) Show() {
	d.showAble.Show()
}

type FoundationMakeUp struct {
	Decorator
}

func NewFoundationMakeUp(showAble ShowAble) *FoundationMakeUp {
	return &FoundationMakeUp{Decorator{showAble}}
}

func (d *FoundationMakeUp) Show() {
	fmt.Print("打粉底【")
	d.Decorator.Show()
	fmt.Print("】")
}

// Lipstick 表示口红装饰器
type Lipstick struct {
	Decorator // 继承自 Decorator 基类
}

// NewLipstick 创建一个新的口红装饰器
func NewLipstick(showable ShowAble) *Lipstick {
	return &Lipstick{Decorator: Decorator{showAble: showable}}
}

// Show 展示涂口红的行为
func (l *Lipstick) Show() {
	fmt.Print("涂口红【")
	l.Decorator.Show()
	fmt.Print("】")
}
