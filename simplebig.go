package simplebig

import (
	"math"
	"math/big"
)

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) Int {
	return Int{
		bigInt: big.NewInt(x),
	}
}

// NewIntFromFloat returns new Int set to int part of x which multiplied by 10**decimals
func NewIntFromFloat(x Float, decimals int) Int {
	simpleInt, _ := x.Mul(NewFloat(math.Pow(10, float64(decimals)))).Int()
	return simpleInt
}

// NewIntFromBigInt allocates and returns new Int with value of x.
func NewIntFromBigInt(x *big.Int) Int {
	return Int{
		bigInt: big.NewInt(0).Set(x),
	}
}

// NewIntFromStringFloat allocates and returns a Int with value of int part of s which multiplied by
// 10**decimals
func NewIntFromStringFloat(s string, decimals int) (Int, error) {
	simpleFloat, err := NewFloatFromString(s, 10)
	if err != nil {
		return NewInt(0), err
	}

	return NewIntFromFloat(simpleFloat, decimals), err
}

// NewIntFromString allocates and returns new Int and a boolean indicating of success.
func NewIntFromString(s string, base int) (Int, bool) {
	simpleInt := NewInt(0)
	_, ok := simpleInt.SetString(s, base)
	return simpleInt, ok
}

// NewFloat allocates and returns a new Float sets to x.
func NewFloat(x float64) Float {
	return Float{
		bigFloat: big.NewFloat(x),
	}
}

// NewFloatFromBigFloat allocates a new Float sets to x.
func NewFloatFromBigFloat(x *big.Float) Float {
	simpleFloat := NewFloat(0)
	simpleFloat.bigFloat.Copy(x)
	return simpleFloat
}

// NewIntFromString allocates and returns new Float and a boolean indicating of success.
func NewFloatFromString(s string, base int) (Float, error) {
	f := NewFloat(0)
	_, _, ok := f.Parse(s, base)
	return f, ok
}
