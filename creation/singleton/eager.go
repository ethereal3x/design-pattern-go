package singleton

// 饿汉式单例模式
var eager *Eager

type Eager struct {
	count int
}

func (e *Eager) Inc() {
	e.count++
}

func InitEager(count int) {
	eager = &Eager{count: count}
}

func GetEager() *Eager {
	return eager
}
