package money

import (
	"strconv"
	"strings"
	"unicode"
)

var (
	DecimalParseErrorInvalidInput = Error("An invalid input was provided.")
	DecimalParseErrorNoInput      = Error("No input was provided.")
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
	i, m, _ := strings.Cut(input, ".")

	vErr := validateInput(i + m)

	if vErr != nil {
		return Decimal{}, vErr
	}

	integer, err := strconv.ParseInt(i+m, 10, 64)

	if err != nil {
		return Decimal{}, err
	}

	return Decimal{integer: integer, precision: len(m)}, nil

	// return Decimal{}, nil
}

func validateInput(joined string) error {
	if len(joined) == 0 {
		return DecimalParseErrorNoInput
	}

	for _, ch := range joined {
		if !unicode.IsNumber(rune(ch)) {
			return DecimalParseErrorInvalidInput
		}
	}

	return nil
}

// func stripSeparators(joined string)  {}
