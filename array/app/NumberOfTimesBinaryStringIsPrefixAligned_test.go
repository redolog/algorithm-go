package app

import "testing"

func Test_numTimesAllBlue(t *testing.T) {
	type args struct {
		flips []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				flips: []int{3, 2, 4, 1, 5},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				flips: []int{4, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				flips: []int{2, 1, 3, 5, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTimesAllBlue(tt.args.flips); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}
