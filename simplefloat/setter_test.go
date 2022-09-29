package simplefloat

import (
	"math/big"
	"reflect"
	"testing"
)

func newPtr(x float64) *Float {
	return &Float{
		bigFloat: big.NewFloat(x),
	}
}

func TestFloat_SetUint64(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z uint64
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
				bigFloat: big.NewFloat(10),
			},
			args: args{
				z: 59,
			},
			want: newPtr(59),
		}, {
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(0),
			},
			args: args{
				z: 0,
			},
			want: newPtr(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.SetUint64(tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.SetUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_SetInt64(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z int64
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
				bigFloat: big.NewFloat(10.5),
			},
			args: args{
				z: 59,
			},
			want: newPtr(59),
		},
		{
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(0),
			},
			args: args{
				z: 0,
			},
			want: newPtr(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.SetInt64(tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.SetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_SetFloat64(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z float64
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
				bigFloat: big.NewFloat(10),
			},
			args: args{
				z: 59.5,
			},
			want: newPtr(59.5),
		}, {
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(0),
			},
			args: args{
				z: 0,
			},
			want: newPtr(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.SetFloat64(tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.SetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_SetInt(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z *big.Int
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
				bigFloat: big.NewFloat(10),
			},
			args: args{
				z: big.NewInt(12),
			},
			want: newPtr(12),
		}, {
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(0),
			},
			args: args{
				z: big.NewInt(0),
			},
			want: newPtr(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.SetInt(tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.SetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat_Set(t *testing.T) {
	type fields struct {
		bigFloat *big.Float
	}
	type args struct {
		z *Float
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
				bigFloat: big.NewFloat(10),
			},
			args: args{
				z: newPtr(1.5),
			},
			want: newPtr(1.5),
		}, {
			name: "2",
			fields: fields{
				bigFloat: big.NewFloat(0),
			},
			args: args{
				z: newPtr(0),
			},
			want: newPtr(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Float{
				bigFloat: tt.fields.bigFloat,
			}
			if got := x.Set(tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
