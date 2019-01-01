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
		// TODO: Add test cases.
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
