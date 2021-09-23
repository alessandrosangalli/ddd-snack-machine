package snack_machine

import (
	"ddd-snack-machine/money"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestTotalMoney(t *testing.T) {
	snackMachine := NewSnackMachine(
		money.OneCent(),
		money.TenCents(),
		money.TwentyDollars(),
	)

	totalMoney := snackMachine.getTotalMoney()
	expected := decimal.NewFromFloat(20.11)

	if !totalMoney.Equal(expected) {
		expectedValue, _ := expected.Float64()
		totalMoneyValue, _ := totalMoney.Float64()
		t.Errorf("Error on total value, expected: %f, current: %f", expectedValue, totalMoneyValue)
	}
}

func TestInsertedMoneyAllTypes(t *testing.T) {
	moneyToInsert := []money.Money{
		money.OneCent(),
		money.TenCents(),
		money.TwentyFiveCents(),
		money.OneDollar(),
		money.FiveDollars(),
		money.TwentyDollars(),
	}

	snackMachine := NewSnackMachine(moneyToInsert...)

	if !reflect.DeepEqual(snackMachine.insertedMoney, moneyToInsert) {
		t.Errorf("Inserted money is not equal snack machine money")
	}
}

func TestInsertedMoneyFiveDollars(t *testing.T) {
	moneyToInsert := []money.Money{
		money.TwentyDollars(),
	}

	snackMachine := NewSnackMachine(moneyToInsert...)

	if !reflect.DeepEqual(snackMachine.insertedMoney, moneyToInsert) {
		t.Errorf("Inserted money is not equal snack machine money")
	}
}

func TestInsertedMoneyEmpty(t *testing.T) {
	moneyToInsert := []money.Money{}

	snackMachine := NewSnackMachine(moneyToInsert...)

	if !reflect.DeepEqual(snackMachine.insertedMoney, moneyToInsert) {
		t.Errorf("Inserted money is not equal snack machine money")
	}
}
