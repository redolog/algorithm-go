package decimal

import "testing"

func Test_pivotInteger(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 8,
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				n: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotInteger(tt.args.n); got != tt.want {
				t.Errorf("pivotInteger() = %v, want %v", got, tt.want)
			}
			if got := pivotInteger2(tt.args.n); got != tt.want {
				t.Errorf("pivotInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arithmeticProgressionSum(t *testing.T) {
	type args struct {
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				l: 1,
				r: 1,
			},
			want: 1 * 2,
		},
		{
			name: "test2",
			args: args{
				l: 1,
				r: 6,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arithmeticProgressionSum(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("arithmeticProgressionSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
