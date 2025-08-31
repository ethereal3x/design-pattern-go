package visitor

import "fmt"

type Visitor interface {
	VisitLeopardSpot(leopard *LeopardSpot)
	VisitDolphinSpot(dolphin *DolphinSpot)
}

type Scenery interface {
	Accept(visitor Visitor)
	Price() int
}

type Zoo struct {
	// 动物园包含多个景点
	Sceneries []Scenery
}

func NewZoo() *Zoo {
	return &Zoo{}
}

func (z *Zoo) Add(scenery Scenery) {
	z.Sceneries = append(z.Sceneries, scenery)
}

func (z *Zoo) Accept(v Visitor) {
	for _, scenery := range z.Sceneries {
		scenery.Accept(v)
	}
}

type LeopardSpot struct{}

func (l *LeopardSpot) Accept(visitor Visitor) {
	visitor.VisitLeopardSpot(l)
}

func (l *LeopardSpot) Price() int {
	return 15
}

type DolphinSpot struct{}

func NewDolphinSpot() *DolphinSpot {
	return &DolphinSpot{}
}

func (d *DolphinSpot) Accept(visitor Visitor) {
	visitor.VisitDolphinSpot(d)
}

func (d *DolphinSpot) Price() int {
	return 15
}

type StudentVisitor struct{}

func NewStudentVisitor() *StudentVisitor {
	return &StudentVisitor{}
}

func (s *StudentVisitor) VisitLeopardSpot(leopard *LeopardSpot) {
	fmt.Printf("学生游客游览豹子馆票价: %v\n", leopard.Price()/2)
}

func (s *StudentVisitor) VisitDolphinSpot(dolphin *DolphinSpot) {
	fmt.Printf("学生游客游览海豚馆票价: %v\n", dolphin.Price()/2)
}

type CommonVisitor struct{}

func NewCommonVisitor() *CommonVisitor {
	return &CommonVisitor{}
}

func (c *CommonVisitor) VisitLeopardSpot(leopard *LeopardSpot) {
	fmt.Printf("普通游客游览豹子馆票价: %v\n", leopard.Price())
}

func (c *CommonVisitor) VisitDolphinSpot(dolphin *DolphinSpot) {
	fmt.Printf("普通游客游览海豚馆票价: %v\n", dolphin.Price())
}
