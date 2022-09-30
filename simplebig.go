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

// NewIntFromBigFloat returns new Int set to x.
func NewIntFromBigFloat(x *big.Float, decimals float64) Int {
	copiedValue := new(big.Float)
	copiedValue.Copy(x)
	copiedValue.Mul(copiedValue, big.NewFloat(math.Pow(10, decimals)))

	result := new(big.Int)
	copiedValue.Int(result)
	return Int{
		bigInt: result,
	}
}

// NewIntFromBigInt allocates and returns new Int with value of x.
func NewIntFromBigInt(x *big.Int) Int {
	return Int{
		bigInt: big.NewInt(0).Set(x),
	}
}

// NewIntFromString allocates and returns new Int and a boolean indicating of success.
func NewIntFromString(s string, base int) (Int, bool) {
	simpleInt := NewInt(0)
	_, ok := simpleInt.SetString(s, base)
	return simpleInt, ok
}

// NewFloat allocates and returns a new Float set to x.
func NewFloat(x float64) Float {
	return Float{
		bigFloat: big.NewFloat(x),
	}
}

// NewIntFromString allocates and returns new Float and a boolean indicating of success.
func NewFloatFromString(s string, base int) (Float, error) {
	f := NewFloat(0)
	_, _, ok := f.Parse(s, base)
	return f, ok
}
