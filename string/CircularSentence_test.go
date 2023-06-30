package string

import "testing"

func Test_isCircularSentence(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				sentence: "leetcode",
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				sentence: "Leetcode is cool",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				sentence: "leetcode exercises sound delightful",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				sentence: "eetcode",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCircularSentence(tt.args.sentence); got != tt.want {
				t.Errorf("isCircularSentence() = %v, want %v", got, tt.want)
			}
			if got := isCircularSentence2(tt.args.sentence); got != tt.want {
				t.Errorf("isCircularSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
