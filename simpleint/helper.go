package simpleint

import "fmt"

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

func (x *Int) Scan(s fmt.ScanState, ch rune) error {
	return x.bigInt.Scan(s, ch)
}
