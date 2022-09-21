package simpleint

import (
	"math/big"
	"math/rand"
)

type Int struct {
	bigInt *big.Int
}

// TODO: add String and Scan

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
func (x Int) Sign() int {
	return x.bigInt.Sign()
}

// Bits provides raw (unchecked but fast) access to x by returning its
// absolute value as a little-endian Word slice. The result and x share
// the same underlying array.
// Bits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func (x *Int) Bits() []big.Word {
	return x.bigInt.Bits()
}

// Abs returns the absolute value of x.
func (x Int) Abs() Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Abs(x.bigInt),
	}
}

// Neg returns -x.
func (x Int) Neg() Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Neg(x.bigInt),
	}
}

// Add returns sum x+y.
func (x Int) Add(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Add(x.bigInt, y.bigInt),
	}
}

// Sub returns the difference x-y.
func (x Int) Sub(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Sub(x.bigInt, y.bigInt),
	}
}

// Mul returns the product x*y.
func (x Int) Mul(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Mul(x.bigInt, y.bigInt),
	}
}

// MulRange returns the product of all integers
// in the range [a, b] inclusively.
// If a > b (empty range), the result is 1.
func (a Int) MulRange(b int64) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.MulRange(a.bigInt.Int64(), b),
	}
}

// Binomial returns the binomial coefficient of (n, k).
func (n Int) Binomial(k int64) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Binomial(n.bigInt.Int64(), k),
	}
}

// Quo returns the quotient x/y for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
// Quo implements truncated division (like Go); see QuoRem for more details.
func (x Int) Quo(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Quo(x.bigInt, y.bigInt),
	}
}

// Rem returns the remainder x%y for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
// Rem implements truncated modulus (like Go); see QuoRem for more details.
func (x Int) Rem(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Rem(x.bigInt, y.bigInt),
	}
}

// QuoRem return the quotient x/y and the remainder x%y
// for more informations, check big.Int QuoRem method
func (x Int) QuoRem(y Int) (Int, Int) {
	return x.Quo(y), x.Rem(y)
}

// Div returns the quotient x/y for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see DivMod for more details.
func (x Int) Div(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Div(x.bigInt, y.bigInt),
	}
}

// Mod returns the modulus x%y for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean modulus (unlike Go); see DivMod for more details.
func (x Int) Mod(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Mod(x.bigInt, y.bigInt),
	}
}

// DivMod returns the quotient x div y and z to the modulus x mod y
// for more informations, go to big.Int
func (x Int) DivMod(y Int) (Int, Int) {
	return x.Div(y), x.Mod(y)
}

// Cmp compares x and y and returns:
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
func (x Int) Cmp(y Int) int {
	return x.bigInt.Cmp(y.bigInt)
}

// CmpAbs compares the absolute values of x and y and returns:
//
//	-1 if |x| <  |y|
//	 0 if |x| == |y|
//	+1 if |x| >  |y|
func (x Int) CmpAbs(y Int) int {
	return x.bigInt.CmpAbs(y.bigInt)
}

// Int64 returns the int64 representation of x.
// If x cannot be represented in an int64, the result is undefined.
func (x Int) Int64() int64 {
	return x.bigInt.Int64()
}

// Uint64 returns the uint64 representation of x.
// If x cannot be represented in a uint64, the result is undefined.
func (x Int) Uint64() uint64 {
	return x.bigInt.Uint64()
}

// IsInt64 reports whether x can be represented as an int64.
func (x Int) IsInt64() bool {
	return x.bigInt.IsInt64()
}

// IsUint64 reports whether x can be represented as a uint64.
func (x Int) IsUint64() bool {
	return x.bigInt.IsUint64()
}

// Bytes returns the absolute value of x as a big-endian byte slice.
//
// To use a fixed length slice, or a preallocated one, use FillBytes.
func (x Int) Bytes() []byte {
	return x.bigInt.Bytes()
}

// FillBytes sets buf to the absolute value of x, storing it as a zero-extended
// big-endian byte slice, and returns buf.
//
// If the absolute value of x doesn't fit in buf, FillBytes will panic.
func (x Int) FillBytes(buf []byte) []byte {
	return x.bigInt.FillBytes(buf)
}

// BitLen returns the length of the absolute value of x in bits.
// The bit length of 0 is 0.
func (x Int) BitLen() int {
	return x.bigInt.BitLen()
}

// TrailingZeroBits returns the number of consecutive least significant zero
// bits of |x|.
func (x Int) TrailingZeroBits() uint {
	return x.bigInt.TrailingZeroBits()
}

// Exp returns x**y mod |m| (i.e. the sign of m is ignored).
// If m == nil or m == 0, x = x**y unless y <= 0 then x = 1. If m != 0, y < 0,
// and y and m are not relatively prime, x is unchanged and nil is returned.
//
// Modular exponentiation of inputs of a particular size is not a
// cryptographically constant-time operation.
func (x Int) Exp(y, m Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Exp(x.bigInt, y.bigInt, m.bigInt),
	}
}

// GCD sets x to the greatest common divisor of a and b and returns z.
// for more informations, check big.Int GCD method.
func (z Int) GCD(x, y, a, b Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.GCD(x.bigInt, y.bigInt, a.bigInt, b.bigInt),
	}
}

// Rand sets x to a pseudo-random number in [0, n) and returns x.
//
// As this uses the math/rand package, it must not be used for
// security-sensitive work. Use crypto/rand.Int instead.
func (x *Int) Rand(rnd *rand.Rand, n *Int) *Int {
	x.bigInt.Rand(rnd, n.bigInt)
	return x
}

// TODO
// ModInverse sets x to the multiplicative inverse of x in the ring ℤ/nℤ
// and returns x. If x and n are not relatively prime, x has no multiplicative
// inverse in the ring ℤ/nℤ.  In this case, x is unchanged and the return value
// is x.
// func (x *Int) ModInverse(n *Int) *Int {
// 	bigInt := new(big.Int)
// 	return &Int{
// 		bigInt: bigInt.ModInverse(x.bigInt, n.bigInt),
// 	}
// }

// Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0.
// The y argument must be an odd integer.
func Jacobi(x, y Int) int {
	return big.Jacobi(x.bigInt, y.bigInt)
}

// TODO
// ModSqrt sets z to a square root of x mod p if such a square root exists, and
// returns z. The modulus p must be an odd prime. If x is not a square mod p,
// ModSqrt leaves z unchanged and returns nil. This function panics if p is
// not an odd integer.
// func (z *Int) ModSqrt(x, p *Int) *Int {
// 	return z
// }

// Lsh returns x << n.
func (x Int) Lsh(n uint) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Lsh(x.bigInt, n),
	}
}

// Rsh returns x >> n.
func (x Int) Rsh(n uint) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Rsh(x.bigInt, n),
	}
}

// Bit returns the value of the i'th bit of x. That is, it
// returns (x>>i)&1. The bit index i must be >= 0.
func (x Int) Bit(i int) uint {
	return x.bigInt.Bit(i)
}

// And returns x & y.
func (x Int) And(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.And(x.bigInt, y.bigInt),
	}
}

// AndNot return x &^ y.
func (x Int) AndNot(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.AndNot(x.bigInt, y.bigInt),
	}
}

// Or returns x | y.
func (x Int) Or(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Or(x.bigInt, y.bigInt),
	}
}

// Xor returns x ^ y.
func (x Int) Xor(y Int) Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Xor(x.bigInt, y.bigInt),
	}
}

// Not returns ^x.
func (x Int) Not() Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Not(x.bigInt),
	}
}

// Sqrt returns ⌊√x⌋, the largest integer such that z² ≤ x,.
// It panics if x is negative.
func (x Int) Sqrt() Int {
	bigInt := new(big.Int)
	return Int{
		bigInt: bigInt.Sqrt(x.bigInt),
	}
}
