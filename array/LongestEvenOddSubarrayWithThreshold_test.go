package array

import "testing"

func Test_longestAlternatingSubarray(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:      []int{3, 2, 5, 4},
				threshold: 5,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				nums:      []int{1, 2},
				threshold: 2,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				nums:      []int{2, 3, 4, 5},
				threshold: 4,
			},
			want: 3,
		},
		{
			name: "test4",
			args: args{
				nums:      []int{4, 10, 3},
				threshold: 10,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestAlternatingSubarray(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("longestAlternatingSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
