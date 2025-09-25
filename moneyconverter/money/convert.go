package money

// converter struct
// A converter needs to take in the exchange rate via DI.
// has convert as a receiver func

// applyExchangeRate is in
// charge of multiplying the input quantity by the change rate and returning an
// Amount compatible with the target Currency. We will first need a function to
// multiply a Decimal with an ExchangeRate, and then we’ll have to adjust the
// precision of that product to match the Currency’s.

type converter struct {
	rateService RateService
}

func NewConverter(rateService RateService) *converter {
	return &converter{rateService: rateService}
}

func (c *converter) Convert(amount Amount, to Currency) (Amount, error) {
	// Get rate first.
	_, err := c.rateService.getRate(to.ISOCode)

	if err != nil {
		return Amount{}, err
	}

	return Amount{}, nil
}

func multiply(lhs Decimal, rhs Decimal) Decimal {
	dec := Decimal{
		integer:   lhs.integer * rhs.integer,
		precision: lhs.precision + rhs.precision,
	}

	dec.simplifyInteger()

	return dec
}
