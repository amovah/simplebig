package simplebig

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
)

func TestInt_StringFloat(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		precision int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(101),
			},
			args: args{
				precision: 1,
			},
			want: "10.1",
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(1000),
			},
			args: args{
				precision: 1,
			},
			want: "100.0",
		},
		{
			name: "3",
			fields: fields{
				bigInt: big.NewInt(1000),
			},
			args: args{
				precision: 2,
			},
			want: "10.0",
		},
		{
			name: "4",
			fields: fields{
				bigInt: big.NewInt(200003),
			},
			args: args{
				precision: 3,
			},
			want: "200.003",
		},
		{
			name: "5",
			fields: fields{
				bigInt: big.NewInt(100),
			},
			args: args{
				precision: 0,
			},
			want: "100.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.StringFloat(tt.args.precision); got != tt.want {
				t.Errorf("Int.StringFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Floater(t *testing.T) {
	type fields struct {
		bigInt *big.Int
	}
	type args struct {
		precision int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "1",
			fields: fields{
				bigInt: big.NewInt(101),
			},
			args: args{
				precision: 1,
			},
			want: "10.1",
		},
		{
			name: "2",
			fields: fields{
				bigInt: big.NewInt(1000),
			},
			args: args{
				precision: 1,
			},
			want: "100",
		},
		{
			name: "3",
			fields: fields{
				bigInt: big.NewInt(1000),
			},
			args: args{
				precision: 2,
			},
			want: "10",
		},
		{
			name: "4",
			fields: fields{
				bigInt: big.NewInt(200003),
			},
			args: args{
				precision: 3,
			},
			want: "200.003",
		},
		{
			name: "5",
			fields: fields{
				bigInt: big.NewInt(100),
			},
			args: args{
				precision: 0,
			},
			want: "100",
		},
		{
			name: "6",
			fields: fields{
				bigInt: big.NewInt(1203),
			},
			args: args{
				precision: 7,
			},
			want: "0.0001203",
		},
		{
			name: "6",
			fields: fields{
				bigInt: big.NewInt(1200),
			},
			args: args{
				precision: 7,
			},
			want: "0.00012",
		},
		{
			name: "7",
			fields: fields{
				bigInt: big.NewInt(9989),
			},
			args: args{
				precision: 4,
			},
			want: "0.9989",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Int{
				bigInt: tt.fields.bigInt,
			}
			if got := x.Floater(tt.args.precision); got != tt.want {
				t.Errorf("Int.StringFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

type Data struct {
	Value Int `json:"value"`
}

func TestUnmarshal(t *testing.T) {
	jsonData := []byte("{ \"value\": 40000000000000000 }")

	d := &Data{}
	json.Unmarshal(jsonData, d)
	fmt.Println(d)
	mData, err := json.Marshal(d)
	fmt.Println(string(mData), err)
}
