package iterator

type Iterator interface {
	HasNext() bool
	Next() string
}

type ConcreteIterator struct {
	index int
	data  []string
}

func NewConcreteIterator(data []string) *ConcreteIterator {
	return &ConcreteIterator{
		index: 0,
		data:  data,
	}
}

func (iterator *ConcreteIterator) HasNext() bool {
	return iterator.index < len(iterator.data)
}

func (iterator *ConcreteIterator) Next() string {
	if !iterator.HasNext() {
		return ""
	}
	value := iterator.data[iterator.index]
	iterator.index++
	return value
}

type Aggregate interface {
	CreateIterator() Iterator
}

type ConcreteAggregate struct {
	data []string
}

func NewConcreteAggregate(data []string) *ConcreteAggregate {
	return &ConcreteAggregate{
		data: data,
	}
}

func (a *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(a.data)
}
