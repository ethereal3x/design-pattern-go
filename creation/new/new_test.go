package new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductGetName(t *testing.T) {
	assertions := assert.New(t)

	product := NewProduct("Laptop", 20.01)
	assertions.Equal("Laptop", product.GetName())
}

func TestProductGetPrice(t *testing.T) {
	assertions := assert.New(t)

	product := NewProduct("Laptop", 20.01)
	assertions.Equal(20.01, product.GetPrice())
}
