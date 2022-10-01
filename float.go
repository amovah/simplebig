package simplebig

import (
	"math/big"
)

// A nonzero finite Float represents a multi-precision floating point number
//
//	sign × mantissa × 2**exponent
//
// with 0.5 <= mantissa < 1.0, and MinExp <= exponent <= MaxExp.
// A Float may also be zero (+0, -0) or infinite (+Inf, -Inf).
// All Floats are ordered, and the ordering of two Floats x and y
// is defined by x.Cmp(y).
//
// Each Float value also has a precision, rounding mode, and accuracy.
// The precision is the maximum number of mantissa bits available to
// represent the value. The rounding mode specifies how a result should
// be rounded to fit into the mantissa bits, and accuracy describes the
// rounding error with respect to the exact result.
//
// Unless specified otherwise, all operations (including setters) that
// specify a *Float variable for the result (usually via the receiver
// with the exception of MantExp), round the numeric result according
// to the precision and rounding mode of the result variable.
//
// If the provided result precision is 0 (see below), it is set to the
// precision of the argument with the largest precision value before any
// rounding takes place, and the rounding mode remains unchanged. Thus,
// uninitialized Floats provided as result arguments will have their
// precision set to a reasonable value determined by the operands, and
// their mode is the zero value for RoundingMode (ToNearestEven).
//
// By setting the desired precision to 24 or 53 and using matching rounding
// mode (typically ToNearestEven), Float operations produce the same results
// as the corresponding float32 or float64 IEEE-754 arithmetic for operands
// that correspond to normal (i.e., not denormal) float32 or float64 numbers.
// Exponent underflow and overflow lead to a 0 or an Infinity for different
// values than IEEE-754 because Float exponents have a much larger range.
//
// The zero (uninitialized) value for a Float is ready to use and represents
// the number +0.0 exactly, with precision 0 and rounding mode ToNearestEven.
//
// Operations always take pointer arguments (*Float) rather
// than Float values, and each unique Float value requires
// its own unique *Float pointer. To "copy" a Float value,
// an existing (or newly allocated) Float must be set to
// a new value using the Float.Set method; shallow copies
// of Floats are not supported and may lead to errors.
type Float struct {
	bigFloat *big.Float
}

// Sign returns:
//
//	-1 if x <   0
//	 0 if x is ±0
//	+1 if x >   0
func (x Float) Sign() int {
	return x.bigFloat.Sign()
}

// MantExp breaks x into its mantissa and exponent components
// and returns the exponent. If a non-nil mant argument is
// provided its value is set to the mantissa of x, with the
// same precision and rounding mode as x. The components
// satisfy x == mant × 2**exp, with 0.5 <= |mant| < 1.0.
// Calling MantExp with a nil argument is an efficient way to
// get the exponent of the receiver.
//
// Special cases are:
//
//	(  ±0).MantExp(mant) = 0, with mant set to   ±0
//	(±Inf).MantExp(mant) = 0, with mant set to ±Inf
//
// x and mant may be the same in which case x is set to its
// mantissa value.
func (x Float) MantExp(mant Float) int {
	return x.bigFloat.MantExp(mant.bigFloat)
}

// SetMantExp returns mant × 2**exp and returns new Float.
// The result has the same precision and rounding mode
// as mant. SetMantExp is an inverse of MantExp but does
// not require 0.5 <= |mant| < 1.0. Specifically, for a
// given x of type *Float, SetMantExp relates to MantExp
// as follows:
//
//	mant := new(Float)
//	new(Float).SetMantExp(mant, x.MantExp(mant)).Cmp(x) == 0
//
// Special cases are:
//
//	x.SetMantExp(  ±0, exp) =   ±0
//	x.SetMantExp(±Inf, exp) = ±Inf
//
// x and mant may be the same in which case x's exponent
// is set to exp.
func (x Float) SetMantExp(mant Float, exp int) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.SetMantExp(mant.bigFloat, exp),
	}
}

// Signbit reports whether x is negative or negative zero.
func (x Float) Signbit() bool {
	return x.bigFloat.Signbit()
}

// IsInf reports whether x is +Inf or -Inf.
func (x Float) IsInf() bool {
	return x.bigFloat.IsInf()
}

// IsInt reports whether x is an integer.
// ±Inf values are not integers.
func (x Float) IsInt() bool {
	return x.bigFloat.IsInt()
}

// Uint64 returns the unsigned integer resulting from truncating x
// towards zero. If 0 <= x <= math.MaxUint64, the result is Exact
// if x is an integer and Below otherwise.
// The result is (0, Above) for x < 0, and (math.MaxUint64, Below)
// for x > math.MaxUint64.
func (x Float) Uint64() (uint64, big.Accuracy) {
	return x.bigFloat.Uint64()
}

// Int64 returns the integer resulting from truncating x towards zero.
// If math.MinInt64 <= x <= math.MaxInt64, the result is Exact if x is
// an integer, and Above (x < 0) or Below (x > 0) otherwise.
// The result is (math.MinInt64, Above) for x < math.MinInt64,
// and (math.MaxInt64, Below) for x > math.MaxInt64.
func (x Float) Int64() (int64, big.Accuracy) {
	return x.bigFloat.Int64()
}

// Float32 returns the float32 value nearest to x. If x is too small to be
// represented by a float32 (|x| < math.SmallestNonzeroFloat32), the result
// is (0, Below) or (-0, Above), respectively, depending on the sign of x.
// If x is too large to be represented by a float32 (|x| > math.MaxFloat32),
// the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.
func (x Float) Float32() (float32, big.Accuracy) {
	return x.bigFloat.Float32()
}

// Float64 returns the float64 value nearest to x. If x is too small to be
// represented by a float64 (|x| < math.SmallestNonzeroFloat64), the result
// is (0, Below) or (-0, Above), respectively, depending on the sign of x.
// If x is too large to be represented by a float64 (|x| > math.MaxFloat64),
// the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.
func (x Float) Float64() (float64, big.Accuracy) {
	return x.bigFloat.Float64()
}

// Int returns the result of truncating x towards zero;
// or nil if x is an infinity.
// The result is Exact if x.IsInt(); otherwise it is Below
// for x > 0, and Above for x < 0.
func (x Float) Int() (Int, big.Accuracy) {
	z := new(big.Int)
	bigInt, acc := x.bigFloat.Int(z)
	return NewIntFromBigInt(bigInt), acc
}

// Copy sets x to y, with the same precision, rounding mode, and
// accuracy as x, and returns x. y is not changed even if x and
// y are the same.
func (x *Float) Copy(y Float) *Float {
	x.bigFloat.Copy(y.bigFloat)
	return x
}

// Abs returns the (possibly rounded) value |x| (the absolute value of x)
func (x Float) Abs() Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Abs(x.bigFloat),
	}
}

// Neg returns the (possibly rounded) value of x with its sign negated,
func (x Float) Neg() Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Neg(x.bigFloat),
	}
}

// Add returns the rounded sum x+y. If x's precision is 0,
// it is changed to the larger of x's or y's precision before the operation.
// Rounding is performed according to x's precision and rounding mode; and
// x's accuracy reports the result error relative to the exact (not rounded)
// result. Add panics with ErrNaN if x and y are infinities with opposite
// signs. The value of z is undefined in that case.
func (x Float) Add(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Add(x.bigFloat, y.bigFloat),
	}
}

// Sub returns the rounded difference x-y.
// Precision, rounding, and accuracy reporting are as for Add.
// Sub panics with ErrNaN if x and y are infinities with equal
// signs. The value of z is undefined in that case.
func (x Float) Sub(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Sub(x.bigFloat, y.bigFloat),
	}
}

// Mul returns the rounded product x*y.
// Precision, rounding, and accuracy reporting are as for Add.
// Mul panics with ErrNaN if one operand is zero and the other
// operand an infinity. The value of z is undefined in that case.
func (x Float) Mul(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Mul(x.bigFloat, y.bigFloat),
	}
}

// Quo returns the rounded quotient x/y.
// Precision, rounding, and accuracy reporting are as for Add.
// Quo panics with ErrNaN if both operands are zero or infinities.
// The value of z is undefined in that case.
func (x Float) Quo(y Float) Float {
	bigFloat := new(big.Float)
	return Float{
		bigFloat: bigFloat.Quo(x.bigFloat, y.bigFloat),
	}
}

// Cmp compares x and y and returns:
//
//	-1 if x <  y
//	 0 if x == y (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
//	+1 if x >  y
func (x Float) Cmp(y Float) int {
	return x.bigFloat.Cmp(y.bigFloat)
}

// Rat returns the rational number corresponding to x;
// or nil if x is an infinity.
// The result is Exact if x is not an Inf.
// If a non-nil *Rat argument z is provided, Rat stores
// the result in z instead of allocating a new Rat.
func (x Float) Rat(z *big.Rat) (*big.Rat, big.Accuracy) {
	return x.bigFloat.Rat(z)
}

// BigFloat returns copy of underlying big.Float
func (x Float) BigFloat() *big.Float {
	bigFloat := new(big.Float)
	bigFloat.Copy(x.bigFloat)
	return bigFloat
}

// Acc returns the accuracy of x produced by the most recent
// operation, unless explicitly documented otherwise by that
// operation.
func (x Float) Acc() big.Accuracy {
	return x.bigFloat.Acc()
}

// MinPrec returns the minimum precision required to represent x exactly
// (i.e., the smallest prec before x.SetPrec(prec) would start rounding x).
// The result is 0 for |x| == 0 and |x| == Inf.
func (x *Float) MinPrec() uint {
	return x.bigFloat.MinPrec()
}

// Mode returns the rounding mode of x.
func (x *Float) Mode() big.RoundingMode {
	return x.bigFloat.Mode()
}
