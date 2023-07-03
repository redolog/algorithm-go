package listNode

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				BuildWithVals([]int{1, 2, 3, 4, 5}),
			},
			want: BuildWithVals([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "test2",
			args: args{
				BuildWithVals([]int{1, 2}),
			},
			want: BuildWithVals([]int{2, 1}),
		},
		{
			name: "test3",
			args: args{
				BuildWithVals([]int{}),
			},
			want: BuildWithVals([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListWithTraverse(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListWithTraverse() = %v, want %v", got, tt.want)
			}
			if got := reverseListWithRecurse(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListWithRecurse() = %v, want %v", got, tt.want)
			}
		})
	}
}
