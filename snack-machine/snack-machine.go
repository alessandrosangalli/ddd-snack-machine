package snack_machine

import (
	"ddd-snack-machine/money"
	"ddd-snack-machine/snack"
	errors "ddd-snack-machine/snack-machine/errors"

	"github.com/shopspring/decimal"
)

type SnackMachine struct {
	totalMoney decimal.Decimal
	snacks     []snack.Snack
}

func NewSnackMachine(inputMoney ...money.Money) *SnackMachine {
	totalMoney := decimal.NewFromInt(0)

	for _, m := range inputMoney {
		totalMoney = decimal.Sum(totalMoney, m.Value)
	}

	snackMachine := SnackMachine{totalMoney: totalMoney}
	return &snackMachine
}

func (snackMachine *SnackMachine) getTotalMoney() decimal.Decimal {
	return snackMachine.totalMoney
}

func (snackMachine *SnackMachine) AddSnack(snack ...snack.Snack) {
	snackMachine.snacks = append(snackMachine.snacks, snack...)
}

func (snackMachine *SnackMachine) Buy(snack snack.Snack) (bool, error) {
	if !snackMachine.snackIsAvaliable(snack) {
		return false, errors.NewInsufficientStock(snack)
	}

	if !snackMachine.getTotalMoney().GreaterThan(snack.Price) {
		return false, errors.NewNoEnoughMoney()
	}

	snackMachine.totalMoney = snackMachine.totalMoney.Sub(snack.Price)

	return true, nil
}

func (snackMachine *SnackMachine) snackIsAvaliable(snack snack.Snack) bool {
	for _, s := range snackMachine.snacks {
		if s.Name == snack.Name {
			return true
		}
	}

	return false
}
