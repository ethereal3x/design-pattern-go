package interpreter

type Expression interface {
	Interpret() int
}

type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

type Add struct {
	left, right Expression
}

func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

type Subtract struct {
	left  Expression
	right Expression
}

func (s *Subtract) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}
