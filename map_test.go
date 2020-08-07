package flatten

import (
	"reflect"
	"testing"
)

func TestMap_Contains(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    Map
		args args
		want bool
	}{
		{
			name: "True/KeyExists",
			m: map[string]string{
				"hello": "world",
			},
			args: args{
				key: "hello",
			},
			want: true,
		},
		{
			name: "False/KeyDoesNotExist",
			m: map[string]string{
				"hello": "world",
			},
			args: args{
				key: "world",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Contains(tt.args.key); got != tt.want {
				t.Errorf("Map.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Delete(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		m    Map
		args args
	}{
		{
			name: "Pass/KeyExistsAndDeletes",
			m: map[string]string{
				"hello": "world",
			},
			args: args{
				prefix: "hello",
			},
		},
		{
			name: "Fail/KeyDoesNotExist",
			m: map[string]string{
				"hello": "world",
			},
			args: args{
				prefix: "world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Delete(tt.args.prefix)
		})
	}
}

func TestMap_Keys(t *testing.T) {
	tests := []struct {
		name string
		m    Map
		want []string
	}{
		{
			name: "Pass/AllKeysReturned",
			m: map[string]string{
				"hello": "world",
			},
			want: []string{"hello"},
		},
		{
			name: "Pass/AllKeysReturned",
			m: map[string]string{
				"hello":   "world",
				"Goodbye": "mars",
				"Moon":    "landing",
			},
			want: []string{"hello", "Goodbye", "Moon"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_Merge(t *testing.T) {
	type args struct {
		m2 Map
	}
	tests := []struct {
		name string
		m    Map
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Merge(tt.args.m2)
		})
	}
}
