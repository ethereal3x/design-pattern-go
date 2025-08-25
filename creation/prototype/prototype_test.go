package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircleClone(t *testing.T) {
	assertions := assert.New(t)

	circle := &Circle{Type: "Circle"}
	cloneCircle := circle.Clone()

	actual := cloneCircle.GetType()
	assertions.Equal(circle.GetType(), actual)
}
