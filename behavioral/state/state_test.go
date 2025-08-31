package state

import "testing"

func TestState(t *testing.T) {
	traffic := NewTrafficLight()
	traffic.ToYellow()
	traffic.ToGreen()
	traffic.ToRed()
}
