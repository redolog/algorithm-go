package twoPointer

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{2, 6, 4, 8, 10, 9, 15},
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
