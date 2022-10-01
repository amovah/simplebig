package simplebig

import (
	"math/big"
	"reflect"
	"testing"
)

func TestInt_Set(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y Int
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
				y: NewInt(15),
			},
			want: newIntPtr(15),
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

func TestInt_SetBit(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		i int
		b uint
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
				bigInt: big.NewInt(23),
			},
			args: args{
				i: 1,
				b: 0,
			},
			want: newIntPtr(big.NewInt(23).SetBit(big.NewInt(23), 1, 0).Int64()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.SetBit(tt.args.i, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetBit() = %v, want %v", got, tt.want)
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
			want: newIntPtr(15),
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
			want: newIntPtr(59),
		}, {
			name: "2",
			fields: fields{
				bigInt: big.NewInt(0),
			},
			args: args{
				y: 0,
			},
			want: newIntPtr(0),
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
			want: newIntPtr(59),
		}, {
			name: "2",
			fields: fields{
				bigInt: big.NewInt(0),
			},
			args: args{
				y: 0,
			},
			want: newIntPtr(0),
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
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(17),
			},
			args: args{
				buf: big.NewInt(50).Bytes(),
			},
			want: newIntPtr(50),
		},
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
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(10),
			},
			args: args{
				s:    "15",
				base: 10,
			},
			want:  newIntPtr(15),
			want1: true,
		},
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(255),
			},
			args: args{
				s:    "1111011",
				base: 2,
			},
			want:  newIntPtr(123),
			want1: true,
		},
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(255),
			},
			args: args{
				s:    "af23",
				base: 2,
			},
			want:  newIntPtr(0),
			want1: false,
		},
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(255),
			},
			args: args{
				s:    "-12",
				base: 10,
			},
			want:  newIntPtr(-12),
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}

			got, got1 := x.SetString(tt.args.s, tt.args.base)
			if got1 != tt.want1 {
				t.Errorf("Int.SetString() got1 = %v, want %v", got1, tt.want1)
			}

			if got1 == false {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetStringFloat(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		s        string
		decimals int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Int
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(10),
			},
			args: args{
				s:        "1001",
				decimals: 1,
			},
			want:    newIntPtr(10010),
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(12),
			},
			args: args{
				s:        "10.1",
				decimals: 1,
			},
			want:    newIntPtr(101),
			wantErr: false,
		},
		{
			name: "3",
			fields: fields{
				bigInt: big.NewInt(51),
			},
			args: args{
				s:        "0.11",
				decimals: 1,
			},
			want:    newIntPtr(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			got, err := x.SetStringFloat(tt.args.s, tt.args.decimals)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int.SetStringFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetStringFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_SetFloat(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		y        Float
		decimals int
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
				y:        NewFloat(50.5),
				decimals: 1,
			},
			want: newIntPtr(505),
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(0),
			},
			args: args{
				y:        NewFloat(5.1),
				decimals: 0,
			},
			want: newIntPtr(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.SetFloat(tt.args.y, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.SetFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
