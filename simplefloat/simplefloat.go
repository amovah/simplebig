package simplefloat

import (
	"math/big"
)

type Float struct {
	bigFloat *big.Float
}

func (x Float) Sign() int {
	return x.bigFloat.Sign()
}

func (x Float) MantExp(mant Float) int {
	return x.bigFloat.MantExp(mant.bigFloat)
}

func (x Float) SetMantExp(mant Float, exp int) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.SetMantExp(mant.bigFloat, exp),
	}
}

func (x Float) Signbit() bool {
	return x.bigFloat.Signbit()
}

func (x Float) IsInf() bool {
	return x.bigFloat.IsInf()
}

func (x Float) IsInt() bool {
	return x.bigFloat.IsInt()
}

func (x Float) Uint64() (uint64, big.Accuracy) {
	return x.bigFloat.Uint64()
}

func (x Float) Int64() (int64, big.Accuracy) {
	return x.bigFloat.Int64()
}

func (x Float) Float32() (float32, big.Accuracy) {
	return x.bigFloat.Float32()
}

func (x Float) Float64() (float64, big.Accuracy) {
	return x.bigFloat.Float64()
}

func (x Float) Int(z *big.Int) (*big.Int, big.Accuracy) {
	return x.bigFloat.Int(z)
}

func (x *Float) Copy(z Float) *Float {
	bigFloat := x.bigFloat.Copy(z.bigFloat)

	return &Float{
		bigFloat: bigFloat,
	}
}

func (x Float) Abs() Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Abs(x.bigFloat),
	}
}

func (x Float) Neg() Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Neg(x.bigFloat),
	}
}

func (x Float) Add(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Add(x.bigFloat, y.bigFloat),
	}
}

func (x Float) Sub(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Sub(x.bigFloat, y.bigFloat),
	}
}

func (x Float) Mul(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Mul(x.bigFloat, y.bigFloat),
	}
}

func (x Float) Quo(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Quo(x.bigFloat, y.bigFloat),
	}
}

func (x Float) Cmp(y Float) int {
	return x.bigFloat.Cmp(y.bigFloat)
}

func (x Float) Rat(z *big.Rat) (*big.Rat, big.Accuracy) {
	return x.bigFloat.Rat(z)
}
