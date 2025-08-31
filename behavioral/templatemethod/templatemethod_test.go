package templatemethod

import (
	"fmt"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	fmt.Println("=======制作红豆豆浆=======")
	redBeanSoyaMilk := NewRedBeanSoyaMilk()
	DoMake(redBeanSoyaMilk)
	fmt.Println("=======制作花生豆浆=======")
	peanutSoyaMilk := NewPeanutSoyaMilk()
	DoMake(peanutSoyaMilk)
}
