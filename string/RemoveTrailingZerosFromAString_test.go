package string

import "testing"

func TestRemoveTrailingZeros(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"51230100"},
			"512301",
		},
		{
			"test2",
			args{"123"},
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeTrailingZeros(tt.args.num); got != tt.want {
				t.Errorf("removeTrailingZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
