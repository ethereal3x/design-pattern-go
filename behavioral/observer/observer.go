package observer

import "fmt"

type Subject interface {
	Register(observer Observer)   // Register 注册一个观察者
	Deregister(observer Observer) // Deregister 注销一个观察者
	Notify(message string)        // Notify 发送通知给所有观察者
}

type Observer interface {
	Update(message string) // Update 接收更新通知
}

type ConcreteSubject struct {
	observers []Observer // 观察者列表
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Deregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(message string) {
	fmt.Println("系统：韭菜们，股票暴涨，大家快买！")
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

type ConcreteObserver struct {
	name string // 观察者名称
}

// NewConcreteObserver 创建一个新的 ConcreteObserver 实例
func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

// Update 实现了 Observer 接口的更新方法
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s: 收到信息<%s>并激进购入股票！\n", o.name, message)
}
