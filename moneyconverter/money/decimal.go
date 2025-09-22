package money

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	errDecimalParseInvalidInput = errors.New("An invalid input was provided.")
	errDecimalParseNoInput      = errors.New("No input was provided.")
)

var (
	DecimalParseInvalidInputError = errDecimalParseInvalidInput
	DecimalParseNoInputError      = errDecimalParseNoInput
)

type Decimal struct {
	integer   int64
	precision int
}

// Takes a given string and converts it to a typical decimal.
// Stores the entire number, and the length of the mantissa as variables in the struct.
// Not sure as of yet why we're doing this, but I'm sure it will be apparent as we progress!
func ParseDecimal(input string) (Decimal, error) {
	// unused variable is presence of separator - will probably see usage in commented out function.
	lhs, rhs, _ := strings.Cut(input, ".")

	vErr := validateDecimalInput(lhs + rhs)

	if vErr != nil {
		return Decimal{}, vErr
	}

	integer, err := strconv.ParseInt(lhs+rhs, 10, 64)

	if err != nil {
		return Decimal{}, err
	}

	res := Decimal{integer: integer, precision: len(rhs)}
	res.simplifyMantissa()

	return res, nil
}

// We don't need trailing zeroes when dealing with precision.
// This function handles that.
func (d *Decimal) simplifyMantissa() {
	// loop through each integer, going backwards
	// if integer == 0, remove it, and decrement the precision property.

	str := strconv.Itoa(int(d.integer))

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == "0" {
			str = str[0:i]
			d.precision--
		}
	}

	simplified, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return
	}

	d.integer = simplified
}

func validateDecimalInput(joined string) error {
	if len(joined) == 0 {
		return errDecimalParseNoInput
	}

	for _, ch := range joined {
		if !unicode.IsNumber(rune(ch)) {
			return errDecimalParseInvalidInput
		}
	}

	return nil
}
