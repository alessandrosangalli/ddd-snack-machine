package snack_machine

import (
	"ddd-snack-machine/money"
	"testing"

	"github.com/shopspring/decimal"
)

func TestInsertMoney(t *testing.T) {
	snackMachine := NewSnackMachine(
		money.OneCent(),
		money.TenCents(),
		money.TwentyDollars(),
	)

	totalMoney := snackMachine.GetTotalMoney()
	expected := decimal.NewFromFloat(20.11)

	if !totalMoney.Equal(expected) {
		expectedValue, _ := expected.Float64()
		totalMoneyValue, _ := totalMoney.Float64()
		t.Errorf("Error on insert value, expected: %f, current: %f", expectedValue, totalMoneyValue)
	}

}
