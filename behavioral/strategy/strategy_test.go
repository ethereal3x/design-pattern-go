package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	operator := NewOperator(&Add{})
	result := operator.Calculate(1, 2)
	fmt.Println("Do add:", result)
}
