package objectpool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObjectPool(t *testing.T) {
	assertions := assert.New(t)

	pool := NewObjectPool(3)

	object1 := pool.AcquireObject()
	assertions.Equal(0, object1.ID)

	object2 := pool.AcquireObject()
	assertions.Equal(1, object2.ID)

	pool.ReleaseObject(object1)
	assertions.Equal(0, object1.ID)

	object3 := pool.AcquireObject()
	assertions.Equal(2, object3.ID)
}
