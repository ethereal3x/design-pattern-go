package proxy

import "testing"

func TestProxy(t *testing.T) {
	foursMarket := NewFoursMarketProxy(NewUser("小明"))
	foursMarket.BuyCar()
}
