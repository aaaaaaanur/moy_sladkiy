package task_06

import (
	"reflect"
	"testing"
)

func TestSwap_21(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "слайс с нечетным количеством элементов",
			args: args{x: []int{5, 6, 53, 1, 90, 39, 2}},
			want: []int{2, 39, 90, 1, 53, 6, 5},
		},
		{
			name: "слайс с четным количеством элементов",
			args: args{x: []int{5, 6, 53, 90, 39, 2}},
			want: []int{2, 39, 90, 53, 6, 5},
		},
		{
			name: "пустой слайс",
			args: args{x: []int{}},
			want: []int{},
		},
		{
			name: "слайс с одним элементом",
			args: args{x: []int{2}},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Swap_2(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Swap_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
