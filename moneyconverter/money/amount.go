package money

import "errors"

type Amount struct {
	quantity Decimal
	currency Currency
}

var errAmountPrecisionMismatch = errors.New("The mantissas of the currency and quantity are not compatible.")
var AmountPrecisionMismatchError = errAmountPrecisionMismatch

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {

	// We can allow for a currency precision smaller than the quantity - we can just pad it.
	if quantity.precision > int(currency.Precision) {
		return Amount{}, errAmountPrecisionMismatch
	}

	return Amount{quantity: quantity, currency: currency}, nil
}
