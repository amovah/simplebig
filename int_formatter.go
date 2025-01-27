package simplebig

import (
	"database/sql/driver"
	"fmt"
	"math"
	"math/big"
	"strings"
)

// Text returns the string representation of x in the given base.
// Base must be between 2 and 62, inclusive. The result uses the
// lower-case letters 'a' to 'z' for digit values 10 to 35, and
// the upper-case letters 'A' to 'Z' for digit values 36 to 61.
// No prefix (such as "0x") is added to the string. If x is a nil
// pointer it returns "<nil>".
func (x Int) Text(base int) string {
	return x.bigInt.Text(base)
}

// String returns the decimal representation of x as generated by
// x.Text(10).
func (x Int) String() string {
	return x.bigInt.String()
}

// StringFloat returns strings representation of x in format of float.
// for example if you a number 101 which it has 1 decimal, you can convert
// it to float string, `x.StringFloat(1) == "10.1"`
func (x Int) StringFloat(decimals int) string {
	decimalPart := NewInt(10).Pow(NewInt(int64(decimals)))
	mod := x.Mod(decimalPart)

	modFloat := new(big.Float)
	modFloat.SetPrec(236)
	modFloat.SetMode(big.ToNearestEven)
	modFloat.SetInt(mod.BigInt())
	modFloat.Mul(modFloat, big.NewFloat(math.Pow(10, float64(decimals)*-1)))
	modPartArr := strings.Split(modFloat.String(), ".")

	var modPart string
	if len(modPartArr) < 2 {
		modPart = "0"
	} else {
		modPart = modPartArr[1]
	}

	integer := x.Div(decimalPart)

	return fmt.Sprintf("%s.%s", integer.String(), modPart)
}

// Floater returns strings representation of x in format of float.
// for example if you a number 101 which it has 1 decimal, you can convert
// it to float string, `x.StringFloat(1) == "10.1"`
// it is like to StringFloat but using strings instead casting to big float
func (x Int) Floater(decimals int) string {
	bigString := x.String()
	if decimals == 0 {
		return bigString
	}

	if len(bigString) < decimals {
		untrun := "0." + strings.Repeat("0", decimals-len(bigString)) + bigString
		return strings.TrimRight(untrun, "0")
	}

	rightPart := strings.TrimRight(bigString[len(bigString)-decimals:], "0")
	leftPart := bigString[0 : len(bigString)-decimals]
	if rightPart == "" {
		return leftPart
	}

	return leftPart + "." + rightPart
}

// MarshalJSON implements the json.Marshaler interface.
func (x *Int) MarshalJSON() ([]byte, error) {
	return x.bigInt.MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (x *Int) UnmarshalJSON(text []byte) error {
	newBig := big.NewInt(0)
	if err := newBig.UnmarshalJSON(text); err != nil {
		return err
	}
	x.bigInt = newBig
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (x *Int) MarshalText() ([]byte, error) {
	return x.bigInt.MarshalText()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (x *Int) UnmarshalText(text []byte) error {
	newBig := big.NewInt(0)
	if err := newBig.UnmarshalText(text); err != nil {
		return err
	}
	x.bigInt = newBig
	return x.bigInt.UnmarshalText(text)
}

// GobDecode implements the gob.GobDecoder interface.
func (x *Int) GobDecode(buf []byte) error {
	return x.bigInt.GobDecode(buf)
}

// GobEncode implements the gob.GobEncoder interface.
func (x *Int) GobEncode() ([]byte, error) {
	return x.bigInt.GobEncode()
}

// Format implements fmt.Formatter. It accepts the formats
// 'b' (binary), 'o' (octal with 0 prefix), 'O' (octal with 0o prefix),
// 'd' (decimal), 'x' (lowercase hexadecimal), and
// 'X' (uppercase hexadecimal).
// Also supported are the full suite of package fmt's format
// flags for integral types, including '+' and ' ' for sign
// control, '#' for leading zero in octal and for hexadecimal,
// a leading "0x" or "0X" for "%#x" and "%#X" respectively,
// specification of minimum digits precision, output field
// width, space or zero padding, and '-' for left or right
// justification.
func (x *Int) Format(s fmt.State, ch rune) {
	x.bigInt.Format(s, ch)
}

// Scan implements sql.Scanner interface.
func (x *Int) Scan(dbValue interface{}) error {
	switch src := dbValue.(type) {
	case nil:
		return nil
	case string:
		if src == "" {
			return nil
		}

		stringValue, ok := dbValue.(string)
		if !ok {
			return fmt.Errorf("Scan: unable to cast type %T into string in BigNumber", src)
		}

		_, ok = x.SetString(stringValue, 10)
		if !ok {
			return fmt.Errorf("Scan: unable to set string to big.int in BigNumber, source: %s", src)
		}

		return nil
	default:
		return fmt.Errorf("Scan: unable to scan type %T into BigNumber", src)
	}
}

// Value implements sql/driver.Valuer interface.
func (x Int) Value() (driver.Value, error) {
	return x.String(), nil
}
