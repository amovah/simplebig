package simpleint

import (
	"math"
	"math/big"
)

// New allocates and returns a new Int set to x.
func New(x int64) Int {
	return Int{
		bigInt: big.NewInt(x),
	}
}

func NewBigFloat(bigFloatValue *big.Float, precision float64) Int {
	copiedValue := new(big.Float)
	copiedValue.Copy(bigFloatValue)
	copiedValue.Mul(copiedValue, big.NewFloat(math.Pow(10, precision)))

	result := new(big.Int)
	copiedValue.Int(result)
	return Int{
		bigInt: result,
	}
}

func NewBigInt(bigIntValue *big.Int) Int {
	return Int{
		bigInt: big.NewInt(0),
	}
}
