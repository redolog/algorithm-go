package sort

import "testing"

func Test_buyChoco(t *testing.T) {
	type args struct {
		prices []int
		money  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				prices: []int{1, 2, 2},
				money:  3,
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				prices: []int{3, 2, 3},
				money:  3,
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				prices: []int{69, 91, 78, 19, 40, 13},
				money:  94,
			},
			want: 62,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buyChocoMinMin2(tt.args.prices, tt.args.money); got != tt.want {
				t.Errorf("buyChocoMinMin2() = %v, want %v", got, tt.want)
			}
			if got := buyChocoSort(tt.args.prices, tt.args.money); got != tt.want {
				t.Errorf("buyChocoSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
