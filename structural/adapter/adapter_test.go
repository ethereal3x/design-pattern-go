package adapter

import "testing"

package adapter

func TestAdapter(t *testing.T) {
	plug := &TwoPinPlugin{}

	threePinSocket := ThreePinPowerSocket{}

	//三孔插头是不能为两针插头充电的,可以试试看
	threePinSocket.ThreePinCharging(plug)

	//只好加一个电源适配器
	powerSocket := NewPowerAdapter(&threePinSocket)

	//再试试能不能充电
	powerSocket.PlugCharging(plug)

	// Output:
	// i can not charge for this type
	// charging for three pin plug
}