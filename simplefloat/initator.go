package simplefloat

import "math/big"

func New(x float64) Float {
	return Float{
		bigFloat: big.NewFloat(x),
	}
}
