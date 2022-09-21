package simpleint

import "math/big"

// New allocates and returns a new Int set to x.
func New(x int64) *Int {
	bigInt := new(big.Int)
	bigInt.SetInt64(x)
	return &Int{
		bigInt: bigInt,
	}
}
