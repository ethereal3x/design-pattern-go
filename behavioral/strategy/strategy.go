package strategy

type IStrategy interface {
	Do(int, int) int
}

type Add struct {
}

func (*Add) Do(a int, b int) int {
	return a + b
}

type Reduce struct{}

func (*Reduce) Do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

func (o *Operator) SetStrategy(strategy IStrategy) {
	o.strategy = strategy
}

func (o *Operator) Calculate(a, b int) int {
	return o.strategy.Do(a, b)
}

func NewOperator(strategy IStrategy) *Operator {
	return &Operator{strategy: strategy}
}
