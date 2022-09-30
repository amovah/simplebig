package simplebig

import (
	"math/big"
)

// SetUint64 sets x to the (possibly rounded) value of y and returns x.
// If x's precision is 0, it is changed to 64 (and rounding will have
// no effect).
func (x *Float) SetUint64(y uint64) *Float {
	x.bigFloat.SetUint64(y)
	return x
}

// SetInt64 sets x to the (possibly rounded) value of y and returns x.
// If x's precision is 0, it is changed to 64 (and rounding will have
// no effect).
func (x *Float) SetInt64(y int64) *Float {
	x.bigFloat.SetInt64(y)
	return x
}

// SetFloat64 sets x to the (possibly rounded) value of y and returns x.
// If x's precision is 0, it is changed to 53 (and rounding will have
// no effect). SetFloat64 panics with ErrNaN if x is a NaN.
func (x *Float) SetFloat64(y float64) *Float {
	x.bigFloat.SetFloat64(y)
	return x
}

// SetInt sets x to the (possibly rounded) value of y and returns x.
// If x's precision is 0, it is changed to the larger of x.BitLen()
// or 64 (and rounding will have no effect).
func (x *Float) SetInt(y *big.Int) *Float {
	x.bigFloat.SetInt(y)
	return x
}

// Set sets x to the (possibly rounded) value of y and returns x.
// If x's precision is 0, it is changed to the precision of y
// before setting x (and rounding will have no effect).
// Rounding is performed according to x's precision and rounding
// mode; and x's accuracy reports the result error relative to the
// exact (not rounded) result.
func (x *Float) Set(y *Float) *Float {
	x.bigFloat.Set(y.bigFloat)
	return x
}

// SetString sets x to the value of s and returns x and a boolean indicating
// success. s must be a floating-point number of the same format as accepted
// by Parse, with base argument 0. The entire string (not just a prefix) must
// be valid for success. If the operation failed, the value of x is undefined
// but the returned value is nil.
func (x *Float) SetString(s string) (*Float, bool) {
	_, bool := x.bigFloat.SetString(s)
	return x, bool
}

// String formats x like x.Text('g', 10).
func (x *Float) String() string {
	return x.bigFloat.String()
}

// SetInf sets x to the infinite Float -Inf if signbit is
// set, or +Inf if signbit is not set, and returns x. The
// precision of x is unchanged and the result is always
// Exact.
func (x *Float) SetInf(signbit bool) *Float {
	x.bigFloat.SetInf(signbit)
	return x
}
