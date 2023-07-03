package listNode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbersII(t *testing.T) {
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
				l1: BuildWithVals([]int{7, 2, 4, 3}),
				l2: BuildWithVals([]int{5, 6, 4}),
			},
			want: BuildWithVals([]int{7, 8, 0, 7}),
		},
		{
			name: "test2",
			args: args{
				l1: BuildWithVals([]int{2, 4, 3}),
				l2: BuildWithVals([]int{5, 6, 4}),
			},
			want: BuildWithVals([]int{8, 0, 7}),
		},
		{
			name: "test3",
			args: args{
				l1: BuildWithVals([]int{0}),
				l2: BuildWithVals([]int{0}),
			},
			want: BuildWithVals([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersII(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbersII() = %v, want %v", got, tt.want)
			}
			if got := addTwoNumbersIIListSolution(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbersIIListSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
