package proxy

import "fmt"

type IBuyCar interface {
	BuyCar()
}

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) BuyCar() {
	fmt.Printf("<%s买车>\n", u.name)
}

type FoursMarketProxy struct {
	user *User
}

func NewFoursMarketProxy(user *User) *FoursMarketProxy {
	return &FoursMarketProxy{user: user}
}

func (f *FoursMarketProxy) BuyCar() {
	fmt.Println("汽车上牌子")
	fmt.Println("注册汽车")
	fmt.Println("从供应商购买汽车到 4S 店")
	f.user.BuyCar()
	fmt.Println("保养汽车")
	fmt.Println("维修汽车")
}
