package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	expression := Subtract{
		left: &Add{
			left:  &Number{value: 4},
			right: &Number{value: 2},
		},
		right: &Number{value: 3},
	}

	result := expression.Interpret()
	fmt.Println(result)
}
