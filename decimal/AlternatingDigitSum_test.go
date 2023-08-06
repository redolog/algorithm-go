package decimal

import "testing"

func Test_alternateDigitSum(t *testing.T) {
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
			args: args{n: 521},
			want: 4,
		},
		{
			name: "test2",
			args: args{n: 111},
			want: 1,
		},
		{
			name: "test3",
			args: args{n: 886996},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternateDigitSum(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSum() = %v, want %v", got, tt.want)
			}
			if got := alternateDigitSum2(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
