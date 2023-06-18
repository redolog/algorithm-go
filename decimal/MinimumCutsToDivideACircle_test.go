package decimal

import "testing"

func Test_numberOfCuts(t *testing.T) {
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
				n: 4,
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				n: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfCuts(tt.args.n); got != tt.want {
				t.Errorf("numberOfCuts() = %v, want %v", got, tt.want)
			}
		})
	}
}
