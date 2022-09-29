package simplefloat

import (
	"math/big"
)

func (x *Float) SetUint64(z uint64) *Float {
	x.bigFloat.SetUint64(z)
	return x
}

func (x *Float) SetInt64(z int64) *Float {
	x.bigFloat.SetInt64(z)
	return x
}

func (x *Float) SetFloat64(z float64) *Float {
	x.bigFloat.SetFloat64(z)
	return x
}

func (x *Float) SetInt(z *big.Int) *Float {
	x.bigFloat.SetInt(z)
	return x
}

func (x *Float) Set(z *Float) *Float {
	x.bigFloat.Set(z.bigFloat)
	return x
}

func (x *Float) SetString(s string) (*Float, bool) {
	bigFloat, bool := x.bigFloat.SetString(s)
	return &Float{bigFloat: bigFloat}, bool
}

func (x *Float) String() string {
	return x.bigFloat.String()
}

func (x *Float) SetInf(signbit bool) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.SetInf(signbit),
	}
}
