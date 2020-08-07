package flatten

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	type args struct {
		thing map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want Map
	}{
		{
			name: "Pass/StringSingleLevel",
			args: args{
				thing: map[string]interface{}{
					"hello": "world",
				},
			},
			want: map[string]string{
				"hello": "world",
			},
		},
		{
			name: "Pass/IntSingleLevel",
			args: args{
				thing: map[string]interface{}{
					"hello": 1234,
				},
			},
			want: map[string]string{
				"hello": "1234",
			},
		},
		{
			name: "Pass/BoolTrueSingleLevel",
			args: args{
				thing: map[string]interface{}{
					"hello": true,
				},
			},
			want: map[string]string{
				"hello": "true",
			},
		},
		{
			name: "Pass/BoolFalseSingleLevel",
			args: args{
				thing: map[string]interface{}{
					"hello": false,
				},
			},
			want: map[string]string{
				"hello": "false",
			},
		},
		{
			name: "Pass/FloatSingleLevel",
			args: args{
				thing: map[string]interface{}{
					"hello": 1234.56,
				},
			},
			want: map[string]string{
				"hello": "1234.560000",
			},
		},
		{
			name: "Pass/SliceSingleElement",
			args: args{
				thing: map[string]interface{}{
					"hello": []string{"world"},
				},
			},
			want: map[string]string{
				"hello.0": "world",
			},
		},
		{
			name: "Pass/SliceManyElement",
			args: args{
				thing: map[string]interface{}{
					"hello": []string{"world", "mars", "tree"},
				},
			},
			want: map[string]string{
				"hello.0": "world",
				"hello.1": "mars",
				"hello.2": "tree",
			},
		},
		{
			name: "Pass/MapSimple",
			args: args{
				thing: map[string]interface{}{
					"hello": map[string]interface{}{
						"world": 27,
					},
				},
			},
			want: map[string]string{
				"hello.world": "27",
			},
		},
		{
			name: "Pass/MapSimpleMany",
			args: args{
				thing: map[string]interface{}{
					"hello": map[string]interface{}{
						"world": 27,
						"mars":  3,
					},
				},
			},
			want: map[string]string{
				"hello.world": "27",
				"hello.mars":  "3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.thing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flatten(t *testing.T) {
	type args struct {
		result map[string]string
		prefix string
		v      reflect.Value
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.result, tt.args.prefix, tt.args.v)
		})
	}
}

func Test_flattenMap(t *testing.T) {
	type args struct {
		result map[string]string
		prefix string
		v      reflect.Value
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flattenMap(tt.args.result, tt.args.prefix, tt.args.v)
		})
	}
}

func Test_flattenSlice(t *testing.T) {
	type args struct {
		result map[string]string
		prefix string
		v      reflect.Value
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flattenSlice(tt.args.result, tt.args.prefix, tt.args.v)
		})
	}
}
