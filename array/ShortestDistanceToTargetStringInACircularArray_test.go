package array

import "testing"

func Test_closetTarget(t *testing.T) {
	type args struct {
		words      []string
		target     string
		startIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				words:      []string{"hello", "i", "am", "leetcode", "hello"},
				target:     "hello",
				startIndex: 1,
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				words:      []string{"a", "b", "leetcode"},
				target:     "leetcode",
				startIndex: 0,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				words:      []string{"i", "eat", "leetcode"},
				target:     "ate",
				startIndex: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closetTarget(tt.args.words, tt.args.target, tt.args.startIndex); got != tt.want {
				t.Errorf("closetTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
