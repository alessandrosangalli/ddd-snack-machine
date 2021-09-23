package snack_machine

import (
	"ddd-snack-machine/money"
	"ddd-snack-machine/snack"
	snackMachineErrors "ddd-snack-machine/snack-machine/errors"
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
		t.Error("Inserted money is not equal snack machine money")
	}
}

func TestBuySnack(t *testing.T) {
	moneyToInsert := []money.Money{}
	snackMachine := NewSnackMachine(moneyToInsert...)
	snackMachine.AddSnack(snack.Chocolate())

	res, err := snackMachine.Buy(snack.Chocolate())
	expected := true

	if res != expected {
		t.Errorf("Error on buy snack, expected: %t, current: %t", expected, res)

		if err != nil {
			t.Errorf("err: %s", err.Error())
		}
	}
}

func TestBuySnackNotAvaliable(t *testing.T) {
	moneyToInsert := []money.Money{}
	snackMachine := NewSnackMachine(moneyToInsert...)

	_, err := snackMachine.Buy(snack.Chocolate())
	insufficientStockChocolate := snackMachineErrors.NewInsufficientStock(snack.Chocolate())

	if err.Error() != insufficientStockChocolate.Error() {
		t.Errorf("Error on buy snack, expected: %s, current: %s", insufficientStockChocolate, err)
	}
}

func TestBuySnackNoEnoughMoney(t *testing.T) {
	moneyToInsert := []money.Money{}
	snackMachine := NewSnackMachine(moneyToInsert...)
	snackMachine.AddSnack(snack.Chocolate())

	_, err := snackMachine.Buy(snack.Chocolate())
	noEnoughMoney := snackMachineErrors.NewNoEnoughMoney()

	if err.Error() != noEnoughMoney.Error() {
		t.Errorf("Error on buy snack, expected: %s, current: %s", noEnoughMoney, err)
	}

}
