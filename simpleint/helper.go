package simpleint

import (
	"database/sql/driver"
	"fmt"
)

func (x Int) Text(base int) string {
	return x.bigInt.Text(base)
}

func (x Int) String() string {
	return x.bigInt.String()
}

func (x *Int) MarshalJSON() ([]byte, error) {
	return x.bigInt.MarshalJSON()
}

func (x *Int) UnmarshalJSON(text []byte) error {
	return x.bigInt.UnmarshalJSON(text)
}

func (x *Int) MarshalText() ([]byte, error) {
	return x.bigInt.MarshalText()
}

func (x *Int) UnmarshalText(text []byte) error {
	return x.bigInt.UnmarshalText(text)
}

func (x *Int) GobDecode(buf []byte) error {
	return x.bigInt.GobDecode(buf)
}

func (x *Int) GobEncode() ([]byte, error) {
	return x.bigInt.GobEncode()
}

func (x *Int) Format(s fmt.State, ch rune) {
	x.bigInt.Format(s, ch)
}

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

func (x Int) Value() (driver.Value, error) {
	return x.String(), nil
}
