package simpleint

import (
	"math/big"
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
