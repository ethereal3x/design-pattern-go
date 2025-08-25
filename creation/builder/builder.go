package builder

import "fmt"

type ICar interface {
	Speed() int
	Brand() string
	Brief()
}

type ICarBuilder interface {
	Wheel(wheel int) ICarBuilder
	Engine(engin string) ICarBuilder
	Speed(max int) ICarBuilder
	Brand(brand string) ICarBuilder
	Build() ICar
}

type CarProto struct {
	Wheel     int
	Engine    string
	MaxSpeed  int
	BrandName string
}

func (c *CarProto) Speed() int {
	return c.MaxSpeed
}

func (c *CarProto) Brand() string {
	return c.BrandName
}

// Brief 简介
func (c *CarProto) Brief() {
	fmt.Println("this is a cool car")
	fmt.Println("car wheel size: ", c.Wheel)
	fmt.Println("car MaxSpeed: ", c.MaxSpeed)
	fmt.Println("car Engine: ", c.Engine)
}

// CarStudio 打算通过成立造车实验室进行造车
type CarStudio struct {
	prototype CarProto
}

// NewCarStudio 造车工作室
func NewCarStudio() ICarBuilder {
	return &CarStudio{}
}

// Wheel of car
func (c *CarStudio) Wheel(wheel int) ICarBuilder {
	c.prototype.Wheel = wheel
	return c
}

// Engine of car
func (c *CarStudio) Engine(engine string) ICarBuilder {
	c.prototype.Engine = engine
	return c
}

// Speed of car
func (c *CarStudio) Speed(max int) ICarBuilder {
	c.prototype.MaxSpeed = max
	return c
}

// Brand of car
func (c *CarStudio) Brand(brand string) ICarBuilder {
	c.prototype.BrandName = brand
	return c
}

// Build return a car
func (c *CarStudio) Build() ICar {
	car := &CarProto{
		Wheel:     c.prototype.Wheel,
		Engine:    c.prototype.Engine,
		MaxSpeed:  c.prototype.MaxSpeed,
		BrandName: c.prototype.BrandName,
	}
	return car
}
