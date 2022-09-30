package simplebig

import "math/big"

// SetBit set i'th bit to b (0 or 1).
// That is, if b is 1 SetBit sets x = y | (1 << i);
// if b is 0 SetBit sets x = y &^ (1 << i). If b is not 0 or 1,
// SetBit will panic.
func (x *Int) SetBit(i int, b uint) *Int {
	x.bigInt.SetBit(x.bigInt, i, b)
	return x
}

// SetBytes interprets buf as the bytes of a big-endian unsigned
// integer, sets x to that value, and returns x.
func (x *Int) SetBytes(buf []byte) *Int {
	x.bigInt.SetBytes(buf)
	return x
}

// Set sets x to y returns x.
func (x *Int) Set(y *Int) *Int {
	x.bigInt.Set(y.bigInt)
	return x
}

// SetBits provides raw (unchecked but fast) access to z by setting its
// value to abs, interpreted as a little-endian Word slice, and returning
// z. The result and abs share the same underlying array.
// SetBits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func (x *Int) SetBits(abs []big.Word) *Int {
	x.bigInt.SetBits(abs)
	return x
}

// SetInt64 sets x to y and returns x.
func (x *Int) SetInt64(y int64) *Int {
	x.bigInt.SetInt64(y)
	return x
}

// SetUint64 sets x to y returns x.
func (x *Int) SetUint64(y uint64) *Int {
	x.bigInt.SetUint64(y)
	return x
}

// SetString sets x to the value of s, interpreted in the given base,
// and returns x and a boolean indicating success. The entire string
// (not just a prefix) must be valid for success.
//
// The base argument must be 0 or a value between 2 and MaxBase.
// For base 0, the number prefix determines the actual base: A prefix of
// “0b” or “0B” selects base 2, “0”, “0o” or “0O” selects base 8,
// and “0x” or “0X” selects base 16. Otherwise, the selected base is 10
// and no prefix is accepted.
//
// For bases <= 36, lower and upper case letters are considered the same:
// The letters 'a' to 'z' and 'A' to 'Z' represent digit values 10 to 35.
// For bases > 36, the upper case letters 'A' to 'Z' represent the digit
// values 36 to 61.
//
// For base 0, an underscore character “_” may appear between a base
// prefix and an adjacent digit, and between successive digits; such
// underscores do not change the value of the number.
// Incorrect placement of underscores is reported as an error if there
// are no other errors. If base != 0, underscores are not recognized
// and act like any other character that is not a valid digit.
func (x *Int) SetString(s string, base int) (*Int, bool) {
	bigInt, bool := x.bigInt.SetString(s, base)
	return &Int{bigInt: bigInt}, bool
}
