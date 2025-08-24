package simplefatory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRMB(t *testing.T) {
	assertions := assert.New(t)
	rmb := NewMoney(MoneyTypeRMB)
	actual := rmb.Cost()
	assertions.Equal("cost rmb", actual)
}
