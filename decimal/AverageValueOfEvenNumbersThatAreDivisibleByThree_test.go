package decimal

import "testing"

func Test_averageValue(t *testing.T) {
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
			args: args{[]int{1, 3, 6, 10, 12, 15}},
			want: 9,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 4, 7, 10}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageValue(tt.args.nums); got != tt.want {
				t.Errorf("averageValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
