package snack_machine

type NoEnoughMoney struct {
	message string
}

func NewNoEnoughMoney() *NoEnoughMoney {
	return &NoEnoughMoney{
		message: "No enough money",
	}
}

func (e *NoEnoughMoney) Error() string {
	return e.message
}
