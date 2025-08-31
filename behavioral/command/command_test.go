package command

import "testing"

func TestCommand(t *testing.T) {
	light := &Light{}
	turnOnCommand := &TurnOnCommand{light: light}
	turnOffCommand := &TurnOffCommand{light: light}
	invoker := Invoker{command: turnOnCommand}
	invoker.ExecuteCommand()
	invoker.command = turnOffCommand
	invoker.ExecuteCommand()
}
