package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	aggregate := NewConcreteAggregate([]string{"apple", "banana", "cherry", "date"})
	iterator := aggregate.CreateIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
