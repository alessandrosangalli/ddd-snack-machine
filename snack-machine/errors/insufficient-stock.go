package snack_machine

import "ddd-snack-machine/snack"

type InsufficientStock struct {
	message string
}

func NewInsufficientStock(snack snack.Snack) *InsufficientStock {
	return &InsufficientStock{
		message: "Insufficient stock of " + snack.Name,
	}
}

func (e *InsufficientStock) Error() string {
	return e.message
}
