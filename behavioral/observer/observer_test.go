package observer

import "testing"

func TestObserver(t *testing.T) {
	// 创建一个主题对象
	subject := NewConcreteSubject()

	// 创建两个观察者对象并注册到主题中
	observer1 := NewConcreteObserver("韭菜一号")
	observer2 := NewConcreteObserver("韭菜二号")
	subject.Register(observer1)
	subject.Register(observer2)

	// 发送通知给所有观察者
	subject.Notify("腾讯股票即将暴涨")

	// 注销观察者 observer2
	subject.Deregister(observer2)

	// 再次发送通知给观察者
	subject.Notify("字节股票即将暴涨")
}
