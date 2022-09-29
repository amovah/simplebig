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

func NewIntFromBigInt(bigInt *big.Int) Int {
	return Int{
		bigInt: big.NewInt(0).Set(bigInt),
	}
}
