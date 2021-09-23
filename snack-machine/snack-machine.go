package snack_machine

import (
	"ddd-snack-machine/money"

	"github.com/shopspring/decimal"
)

type SnackMachine struct {
	insertedMoney []money.Money
}

func NewSnackMachine(inputMoney ...money.Money) *SnackMachine {
	snackMachine := SnackMachine{insertedMoney: inputMoney}
	return &snackMachine
}

func (snackMachine *SnackMachine) getTotalMoney() decimal.Decimal {
	total := decimal.NewFromInt(0)

	for _, m := range snackMachine.insertedMoney {
		total = decimal.Sum(total, m.Value)
	}

	return total
}
