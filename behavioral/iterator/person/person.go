package person

import "fmt"

type Iterator struct {
	index int
	Container
}

func (i *Iterator) Next() Visitor {
	fmt.Print(i.index)
	visitor := i.list[i.index]
	i.index += 1
	return visitor
}

func (i *Iterator) HasNext() bool {
	if i.index >= len(i.list) {
		return false
	}
	return true
}

type Container struct {
	list []Visitor
}

func (c *Container) Add(visitor Visitor) {
	c.list = append(c.list, visitor)
}

func (c *Container) Remove(index int) {
	if index < 0 || index > len(c.list) {
		return
	}
	c.list = append(c.list[:index], c.list[index+1:]...)
}

type Visitor interface {
	Visit()
}

type Teacher struct {
}

type Analysis struct {
}

func (t *Teacher) Visit() {
	fmt.Println("This is teacher visitor")
}

func (a *Analysis) Visit() {
	fmt.Println("This is analyzer visitor")
}

func NewIterator() *Iterator {
	return &Iterator{
		index:     0,
		Container: Container{},
	}
}
