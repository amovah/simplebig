package simpleint

import "math/big"

// New allocates and returns a new Int set to x.
func New(x int64) Int {
	return Int{
		bigInt: big.NewInt(x),
	}
}
