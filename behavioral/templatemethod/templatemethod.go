package templatemethod

import "fmt"

type Milker interface {
	SelectBean()
	Soak()
	Beat()
	AddCondiment()
}

type SoyaMilk struct {
}

func (s *SoyaMilk) SelectBean() {
	fmt.Println("select fresh beans")
}

func (s *SoyaMilk) Soak() {
	fmt.Println("soak bean")
}

func (s *SoyaMilk) Beat() {
	fmt.Println("beat bean")
}

type RedBeanSoyaMilk struct {
	SoyaMilk
}

func NewRedBeanSoyaMilk() *RedBeanSoyaMilk {
	return &RedBeanSoyaMilk{}
}

func (r *RedBeanSoyaMilk) AddCondiment() {
	fmt.Println("add red bean condiment")
}

type PeanutBeanSoyaMilk struct {
	SoyaMilk
}

func NewPeanutSoyaMilk() *PeanutBeanSoyaMilk {
	return &PeanutBeanSoyaMilk{}
}

func (p *PeanutBeanSoyaMilk) AddCondiment() {
	fmt.Println("add peanut bean condiment")
}

func DoMake(milk Milker) {
	milk.SelectBean()
	milk.AddCondiment()
	milk.Soak()
	milk.Beat()
}
