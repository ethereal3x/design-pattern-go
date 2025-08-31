package chainofresponsibility

import "fmt"

type IApprove interface {
	SetNext(IApprove) IApprove
	Approve(amount float64)
}

type Approver struct {
	limit float64
	next  IApprove
}

func NewApprove(limit float64) *Approver {
	return &Approver{limit: limit}
}

func (a Approver) SetNext(next IApprove) IApprove {
	a.next = next
	return a
}

func (a Approver) Approve(amount float64) {

}

type Manager struct {
	*Approver
}

func NewManager(limit float64) *Manager {
	return &Manager{NewApprove(limit)}
}

func (m *Manager) Approve(amount float64) {
	if amount <= m.limit {
		fmt.Printf("Manager approved the request for $%.2f\n", amount)
	} else if m.next != nil {
		m.next.Approve(amount)
	} else {
		fmt.Println("Request can't be approved at Manager level")
	}
}

type Director struct {
	*Approver
}

func NewDirector(limit float64) *Director {
	return &Director{NewApprove(limit)}
}

func (d *Director) Approve(amount float64) {
	if amount <= d.limit {
		fmt.Printf("Director approved the request for $%.2f\n", amount)
	} else if d.next != nil {
		d.next.Approve(amount)
	} else {
		fmt.Println("Request can't be approved at Director level")
	}
}

// CFO 结构体表示CFO职级的审批人
type CFO struct {
	*Approver
}

// NewCFO 创建一个新的CFO审批人
func NewCFO(limit float64) *CFO {
	return &CFO{Approver: NewApprove(limit)}
}

// Approve 实现CFO审批金额的逻辑
func (c *CFO) Approve(amount float64) {
	if amount <= c.limit {
		fmt.Printf("CFO approved the request for $%.2f\n", amount)
	} else {
		fmt.Println("Request can't be approved at CFO level")
	}
}
