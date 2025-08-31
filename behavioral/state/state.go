package state

import "fmt"

type State interface {
	ToGreen(light *TrafficLight)
	ToYellow(light *TrafficLight)
	ToRed(light *TrafficLight)
}

type TrafficLight struct {
	state State
}

func (t *TrafficLight) ToGreen() {
	t.state.ToGreen(t)
}

func (t *TrafficLight) ToYellow() {
	t.state.ToYellow(t)
}

func (t *TrafficLight) ToRed() {
	t.state.ToRed(t)
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{
		state: NewRed(),
	}
}

type Red struct{}

func NewRed() *Red {
	return &Red{}
}

func (r *Red) ToGreen(light *TrafficLight) {
	fmt.Println("错误，红灯不可以切换为绿灯！")
}

func (t *TrafficLight) SetState(state State) {
	t.state = state
}

func (r *Red) ToYellow(light *TrafficLight) {
	fmt.Println("黄灯亮起 5 秒！")
	light.SetState(NewYellow())
}

func (r *Red) ToRed(light *TrafficLight) {
	fmt.Println("错误，已经是红灯！")
}

type Yellow struct{}

func NewYellow() *Yellow {
	return &Yellow{}
}

func (y Yellow) ToGreen(light *TrafficLight) {
	fmt.Println("绿灯亮起 5 秒！")
	light.SetState(NewGreen())
}

func (y Yellow) ToYellow(light *TrafficLight) {
	fmt.Println("错误，已经是黄灯！")
}

func (y Yellow) ToRed(light *TrafficLight) {
	fmt.Println("红灯亮起 5 秒！")
	light.SetState(NewRed())
}

type Green struct{}

func NewGreen() *Green {
	return &Green{}
}

func (g *Green) ToGreen(light *TrafficLight) {
	fmt.Println("错误，已经是绿灯！")
}

func (g *Green) ToYellow(light *TrafficLight) {
	fmt.Println("黄灯亮起 5 秒！")
	light.SetState(NewYellow())
}

func (g *Green) ToRed(light *TrafficLight) {
	fmt.Println("红灯亮起 5 秒！")
	light.SetState(NewRed())
}
