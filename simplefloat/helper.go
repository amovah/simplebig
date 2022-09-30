package simplefloat

import (
	"database/sql/driver"
	"fmt"
	"math"
	"math/big"
	"simplebig/simpleint"
)

func (x Float) Acc() big.Accuracy {
	return x.bigFloat.Acc()
}

func (x Float) Append(buf []byte, fmt byte, prec int) []byte {
	return x.bigFloat.Append(buf, fmt, prec)
}

func (x Float) Format(s fmt.State, format rune) {
	x.bigFloat.Format(s, format)
}

func (x *Float) GobDecode(buf []byte) error {
	return x.bigFloat.GobDecode(buf)
}

func (x *Float) GobEncode() ([]byte, error) {
	return x.bigFloat.GobEncode()

}

func (x *Float) MarshalText() ([]byte, error) {
	return x.bigFloat.MarshalText()
}

func (x *Float) MinPrec() uint {
	return x.bigFloat.MinPrec()
}

func (x *Float) Mode() big.RoundingMode {
	return x.bigFloat.Mode()
}

func (x Float) Parse(s string, base int) (Float, int, error) {
	parsed, b, err := x.bigFloat.Parse(s, base)
	return Float{
		bigFloat: parsed,
	}, b, err
}

func ParseBigFloat(value string) (Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	simpleFloat := New(0)
	simpleFloat.bigFloat = f
	return simpleFloat, err
}

func FloatToBigInt(bigFloatValue *Float, precision float64) simpleint.Int {
	result := new(big.Int)
	copiedValue := new(big.Float)
	copiedValue.Copy(bigFloatValue.bigFloat)
	copiedValue.Mul(copiedValue, big.NewFloat(math.Pow(10, precision)))
	copiedValue.Int(result)
	return simpleint.New(result.Int64())
}

func (x *Float) Scan(dbValue interface{}) error {
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

		_, ok = x.SetString(stringValue)
		if !ok {
			return fmt.Errorf("Scan: unable to set string to big.int in BigNumber, source: %s", src)
		}

		return nil
	default:
		return fmt.Errorf("Scan: unable to scan type %T into BigNumber", src)
	}
}

// Value implements sql/driver.Valuer interface.
func (x Float) Value() (driver.Value, error) {
	return x.String(), nil
}
