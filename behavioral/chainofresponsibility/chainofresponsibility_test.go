package chainofresponsibility

import (
	"fmt"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	manager := NewManager(1000)
	director := NewDirector(5000)
	cfo := NewCFO(10000)

	// 设置审批链
	manager.SetNext(director)
	director.SetNext(cfo)

	// 测试审批流程
	amounts := []float64{800, 3500, 6000, 12000}
	for _, amount := range amounts {
		fmt.Printf("Processing approval request for $%.2f\n", amount)
		manager.Approve(amount)
		fmt.Println()
	}
}
