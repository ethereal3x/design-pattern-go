package simplefatory

type MoneyType string

const (
	MoneyTypeDollar MoneyType = "$"
	MoneyTypeRMB    MoneyType = "￥"
)

type Money interface {
	Cost() string
}

type Dollar struct {
}

type RMB struct {
}

func (*Dollar) Cost() string {
	return "cost dollar"
}

func (*RMB) Cost() string {
	return "cost rmb"
}

func NewMoney(moneyType MoneyType) Money {
	switch moneyType {
	case MoneyTypeDollar:
		return &Dollar{}
	case MoneyTypeRMB:
		return &RMB{}
	default:
		return nil
	}
}
