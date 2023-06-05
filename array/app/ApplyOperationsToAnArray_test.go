package app

import (
	"reflect"
	"testing"
)

func Test_applyOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{nums: []int{1, 2, 2, 1, 1, 0}},
			want: []int{1, 4, 2, 0, 0, 0},
		},
		{
			name: "test2",
			args: args{nums: []int{0, 1}},
			want: []int{1, 0},
		},
		{
			name: "test3",
			args: args{nums: []int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272}},
			want: []int{1694, 399, 832, 1758, 412, 206, 272, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOperations(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
			if got := applyOperations2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
