package money

type ExchangeRate = Decimal

type RateService interface {
	getRate(string) (ExchangeRate, error)
}

type exchangeRateService struct{}

func (ers *exchangeRateService) getRate(string) (ExchangeRate, error) {
	// Hardcode this as 2 for now.
	return ExchangeRate{integer: 2, precision: 0}, nil
}

func NewExchangeRateService() *exchangeRateService {
	return &exchangeRateService{}
}
