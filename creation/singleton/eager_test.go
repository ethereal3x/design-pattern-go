package singleton

import (
	assert "github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// 初始化单例实例
func TestMain(m *testing.M) {
	// 程序运行时初始化
	InitEager(3)
	os.Exit(m.Run())
}

func TestGetEager(t *testing.T) {
	assertions := assert.New(t)
	eager = &Eager{
		count: 3,
	}
	ins := GetEager()
	assertions.Equal(ins, eager)
}
