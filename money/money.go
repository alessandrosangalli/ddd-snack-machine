package money

import "github.com/shopspring/decimal"

type Money struct {
	Value decimal.Decimal
}

func OneCent() Money {
	money := Money{Value: decimal.NewFromFloat(0.01)}

	return money
}

func TenCents() Money {
	money := Money{Value: decimal.NewFromFloat(0.10)}

	return money
}

func TwentyFiveCents() Money {
	money := Money{Value: decimal.NewFromFloat(0.25)}

	return money
}

func OneDollar() Money {
	money := Money{Value: decimal.NewFromInt(1)}

	return money
}

func FiveDollars() Money {
	money := Money{Value: decimal.NewFromInt(5)}

	return money
}

func TwentyDollars() Money {
	money := Money{Value: decimal.NewFromInt(20)}

	return money
}
