package sort

import "testing"

func Test_distinctAverages(t *testing.T) {
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
				[]int{4, 1, 4, 0, 3, 5},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				[]int{1, 100},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				[]int{9, 5, 7, 8, 7, 9, 8, 2, 0, 7},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctAverages(tt.args.nums); got != tt.want {
				t.Errorf("distinctAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
