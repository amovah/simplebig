package simplebig

import (
	"math/big"
	"reflect"
	"testing"
)

func TestInt_New(t *testing.T) {
	type args struct {
		x int64
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{
			name: "1",
			args: args{
				x: 1,
			},
			want: Int{
				bigInt: big.NewInt(1),
			},
		},
		{
			name: "2",
			args: args{
				x: 200,
			},
			want: Int{
				bigInt: big.NewInt(200),
			},
		},
		{
			name: "3",
			args: args{
				x: -32,
			},
			want: Int{
				bigInt: big.NewInt(-32),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInt(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_NewIntFromStringFloat(t *testing.T) {
	type args struct {
		x       string
		decimal int
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{
			name: "1",
			args: args{
				x:       "4999999500.0049849",
				decimal: 7,
			},
			want: Int{
				bigInt: big.NewInt(49999995000049849),
			},
		},
		{
			name: "2",
			args: args{
				x:       "1.03",
				decimal: 7,
			},
			want: Int{
				bigInt: big.NewInt(10300000),
			},
		},
		{
			name: "3",
			args: args{
				x:       "5",
				decimal: 18,
			},
			want: Int{
				bigInt: big.NewInt(5000000000000000000),
			},
		},
		{
			name: "4",
			args: args{
				x:       "0.000000000000001",
				decimal: 7,
			},
			want: Int{
				bigInt: big.NewInt(0),
			},
		},
		{
			name: "5",
			args: args{
				x:       "6.66",
				decimal: 1,
			},
			want: Int{
				bigInt: big.NewInt(66),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := NewIntFromStringFloat(tt.args.x, tt.args.decimal); !reflect.DeepEqual(got, tt.want) && err == nil {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
