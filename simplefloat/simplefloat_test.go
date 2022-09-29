package simplefloat

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFloat_Abs(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Float
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(8.5),
			},
			want: New(8.5),
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(-17.5),
			},
			want: New(17.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Abs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Neg(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Float
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(-8.5),
			},
			want: New(8.5),
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(12.25),
			},
			want: New(-12.25),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Neg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Add(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		y Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Float
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(12.25),
			},
			args: args{
				y: New(5.5),
			},
			want: New(17.75),
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(20),
			},
			args: args{
				y: New(7),
			},
			want: New(27),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Add(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Sub(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		y Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Float
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(12.25),
			},
			args: args{
				y: New(5.5),
			},
			want: New(6.75),
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(20),
			},
			args: args{
				y: New(7),
			},
			want: New(13),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Sub(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Mul(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		y Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Float
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(12.25),
			},
			args: args{
				y: New(5.5),
			},
			want: New(67.375),
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(20),
			},
			args: args{
				y: New(7),
			},
			want: New(140),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Mul(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Quo(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		y Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Float
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(12.25),
			},
			args: args{
				y: New(5),
			},
			want: Float{
				bigFloat: big.NewFloat(0).Quo(big.NewFloat(12.25), big.NewFloat(5)),
			},
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(20),
			},
			args: args{
				y: New(7),
			},
			want: Float{
				bigFloat: big.NewFloat(0).Quo(big.NewFloat(20), big.NewFloat(7)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Quo(tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Quo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Cmp(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		y Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(12.25),
			},
			args: args{
				y: New(5),
			},
			want: 1,
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(0),
			},
			args: args{
				y: New(7),
			},
			want: -1,
		},
		{
			name: "3",
			fields: fields{
				bigFloat: big.NewFloat(10.5),
			},
			args: args{
				y: New(10.5),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Cmp(tt.args.y); got != tt.want {
				t.Errorf("Float.Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Copy(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z Float
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Float
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(1),
			},
			args: args{
				z: New(5),
			},
			want: newPtr(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Copy(tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Sign(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "1",
			fields: fields{
				bigFloat: big.NewFloat(5.5),
			},
			want: 1,
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(0),
			},
			want: 0,
		},
		{
			name: "3",
			fields: fields{
				bigFloat: big.NewFloat(-3.25),
			},
			want: -1,
		},
		{
			name: "4",
			fields: fields{
				bigFloat: big.NewFloat(123.19),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Sign(); got != tt.want {
				t.Errorf("Float.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
