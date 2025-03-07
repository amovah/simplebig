package simplebig

import (
	"fmt"
	"math"
	"math/big"
	"strings"
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
	splitted := strings.Split(s, ".")

	floatPart := ""
	if len(splitted) > 1 {
		floatPart = splitted[1]

		if len(floatPart) > decimals {
			floatPart = floatPart[0:decimals]
		}
	}
	if len(floatPart) < decimals {
		floatPart = floatPart + strings.Repeat("0", decimals-len(floatPart))
	}

	newBigInt, ok := NewIntFromString(splitted[0]+floatPart, 10)
	var err error = nil
	if !ok {
		err = fmt.Errorf("unable to parse %s to simplebig.Int", s)
	}

	return newBigInt, err
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
