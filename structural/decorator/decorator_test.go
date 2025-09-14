package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator_Show(t *testing.T) {
	decoratorGirl := NewLipstick(NewFoundationMakeUp(NewGirl()))
	decoratorGirl.Show()
	fmt.Println()
}
