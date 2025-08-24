package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	// 创建一个木门工厂
	woodenFactory := &WoodenDoorFactory{}

	// 使用木门工厂创建门和门把手
	door := woodenFactory.CreateDoor()
	doorHandle := woodenFactory.CreateDoorHandle()

	// 使用创建的门和门把手
	door.Open()
	doorHandle.Press()
}
