package command

import "fmt"

type Command interface {
	Execute()
}

type Light struct {
}

func (l *Light) TurnOn() {
	fmt.Print("Light is on")
}

func (l *Light) TurnOff() {
	fmt.Print("Light is off")
}

type TurnOnCommand struct {
	light *Light
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

// Invoker 表示调用者
type Invoker struct {
	command Command
}

// ExecuteCommand 调用执行命令
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
