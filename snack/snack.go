package snack

import "github.com/shopspring/decimal"

type Snack struct {
	Price decimal.Decimal
	Name  string
}

func Chocolate() Snack {
	chocolate := Snack{Price: decimal.NewFromInt(3), Name: "Chocolate"}
	return chocolate
}
