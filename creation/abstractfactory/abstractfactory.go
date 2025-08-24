package abstractfactory

import "fmt"

type DoorFactory interface {
	CreateDoor() Door
	CreateDoorHandler() DoorHandle
}

type Door interface {
	Open()
	Close()
}

type DoorHandle interface {
	Press()
}

// WoodenDoorFactory 是一个具体的木门工厂，实现了 DoorFactory 接口
type WoodenDoorFactory struct{}

func (f *WoodenDoorFactory) CreateDoor() Door {
	return &WoodenDoor{}
}

func (f *WoodenDoorFactory) CreateDoorHandle() DoorHandle {
	return &WoodenDoorHandle{}
}

// WoodenDoor 是木门实现
type WoodenDoor struct{}

func (d *WoodenDoor) Open() {
	fmt.Println("Wooden door is opened")
}

func (d *WoodenDoor) Close() {
	fmt.Println("Wooden door is closed")
}

// WoodenDoorHandle 是木门把手实现
type WoodenDoorHandle struct{}

func (h *WoodenDoorHandle) Press() {
	fmt.Println("Press wooden door handle")
}
