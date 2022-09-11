package simpleint

import (
	"math/big"
	"math/rand"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		x int64
	}
	tests := []struct {
		name string
		args args
		want *Int
	}{
		{
			name: "1",
			args: args{
				x: 1,
			},
			want: &Int{
				bigInt: big.NewInt(1),
			},
		},
		{
			name: "2",
			args: args{
				x: 200,
			},
			want: &Int{
				bigInt: big.NewInt(200),
			},
		},
		{
			name: "3",
			args: args{
				x: -32,
			},
			want: &Int{
				bigInt: big.NewInt(-32),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Sign(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(5),
			},
			want: 1,
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(0),
			},
			want: 0,
		},
		{
			name: "3",
			fields: fields{
				bigInt: big.NewInt(-3),
			},
			want: -1,
		},
		{
			name: "4",
			fields: fields{
				bigInt: big.NewInt(123),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Sign(); got != tt.want {
				t.Errorf("Int.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetInt64(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(10),
			},
			args: args{
				y: 59,
			},
			want: New(59),
		}, {
			name: "2",
			fields: fields{
				bigInt: big.NewInt(0),
			},
			args: args{
				y: 0,
			},
			want: New(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.SetInt64(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetUint64(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(10),
			},
			args: args{
				y: 59,
			},
			want: New(59),
		}, {
			name: "2",
			fields: fields{
				bigInt: big.NewInt(0),
			},
			args: args{
				y: 0,
			},
			want: New(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.SetUint64(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Set(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(5),
			},
			args: args{
				y: New(15),
			},
			want: New(15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Set(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Bits(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   []big.Word
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(8),
			},
			want: big.NewInt(8).Bits(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Bits(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Bits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetBits(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		abs []big.Word
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(10),
			},
			args: args{
				abs: big.NewInt(15).Bits(),
			},
			want: New(15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.SetBits(tt.args.abs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Abs(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(8),
			},
			want: New(8),
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(-17),
			},
			want: New(17),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Abs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Neg(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(8),
			},
			want: New(-8),
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(-17),
			},
			want: New(17),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Neg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Add(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(5),
			},
			args: args{
				y: New(7),
			},
			want: New(12),
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(10),
			},
			args: args{
				y: New(-9),
			},
			want: New(1),
		},
		{

			name: "3",
			fields: fields{
				bigInt: big.NewInt(-10),
			},
			args: args{
				y: New(9),
			},
			want: New(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Add(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Sub(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(8),
			},
			args: args{
				y: New(3),
			},
			want: New(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Sub(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Mul(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(10),
			},
			args: args{
				y: New(5),
			},
			want: New(50),
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(5),
			},
			args: args{
				y: New(-1),
			},
			want: New(-5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Mul(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_MulRange(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		b int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(1),
			},
			args: args{
				b: 3,
			},
			want: New(6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := a.MulRange(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.MulRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Binomial(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		k int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(1),
			},
			args: args{
				k: 5,
			},
			want: New(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := n.Binomial(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Binomial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Quo(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Quo(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Quo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Rem(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Rem(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Rem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_QuoRem(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
		z *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
		want1  *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			got, got1 := x.QuoRem(tt.args.y, tt.args.z)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.QuoRem() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Int.QuoRem() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInt_Div(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Div(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Mod(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Mod(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_DivMod(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
		z *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
		want1  *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			got, got1 := x.DivMod(tt.args.y, tt.args.z)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.DivMod() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Int.DivMod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInt_Cmp(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantR  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if gotR := x.Cmp(tt.args.y); gotR != tt.wantR {
				t.Errorf("Int.Cmp() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInt_CmpAbs(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.CmpAbs(tt.args.y); got != tt.want {
				t.Errorf("Int.CmpAbs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Int64(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Int64(); got != tt.want {
				t.Errorf("Int.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Uint64(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Uint64(); got != tt.want {
				t.Errorf("Int.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_IsInt64(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.IsInt64(); got != tt.want {
				t.Errorf("Int.IsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_IsUint64(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.IsUint64(); got != tt.want {
				t.Errorf("Int.IsUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetString(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		s    string
		base int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			got, got1 := x.SetString(tt.args.s, tt.args.base)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Int.SetString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInt_SetBytes(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.SetBytes(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Bytes(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_FillBytes(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.FillBytes(tt.args.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.FillBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_BitLen(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.BitLen(); got != tt.want {
				t.Errorf("Int.BitLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_TrailingZeroBits(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.TrailingZeroBits(); got != tt.want {
				t.Errorf("Int.TrailingZeroBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Exp(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
		z *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Exp(tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_GCD(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
		a *Int
		b *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.GCD(tt.args.y, tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Rand(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		rnd *rand.Rand
		n   *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Rand(tt.args.rnd, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Rand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_ModInverse(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.ModInverse(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.ModInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJacobi(t *testing.T) {
	type args struct {
		x *Int
		y *Int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jacobi(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Jacobi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_ModSqrt(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		x *Int
		p *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := z.ModSqrt(tt.args.x, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.ModSqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Lsh(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		n uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Lsh(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Lsh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Rsh(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		x *Int
		n uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := z.Rsh(tt.args.x, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Rsh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Bit(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Bit(tt.args.i); got != tt.want {
				t.Errorf("Int.Bit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetBit(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
		i int
		b uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.SetBit(tt.args.y, tt.args.i, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_And(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.And(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_AndNot(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.AndNot(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.AndNot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Or(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Or(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Xor(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y *Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Xor(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Not(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Not(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Not() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Sqrt(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Sqrt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
