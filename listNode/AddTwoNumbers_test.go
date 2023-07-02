package listNode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				l1: BuildWithVals([]int{2, 4, 3}),
				l2: BuildWithVals([]int{5, 6, 4}),
			},
			want: BuildWithVals([]int{7, 0, 8}),
		},
		{
			name: "test2",
			args: args{
				l1: BuildWithVals([]int{0}),
				l2: BuildWithVals([]int{0}),
			},
			want: BuildWithVals([]int{0}),
		},
		{
			name: "test3",
			args: args{
				l1: BuildWithVals([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: BuildWithVals([]int{9, 9, 9, 9}),
			},
			want: BuildWithVals([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
