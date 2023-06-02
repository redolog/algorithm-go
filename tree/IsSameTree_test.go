package tree

import (
	"algorithm-go/util"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{
			BuildTreeWithIntArr([]int{1, 2, 3}), BuildTreeWithIntArr([]int{1, 2, 3}),
		}, true},
		{"test2", args{
			BuildTreeWithIntArr([]int{1, 2, 3}), BuildTreeWithQueueNillable([]*int{util.IntPtr(1), util.IntPtr(2), util.IntPtr(3)}),
		}, true},
		{"test3", args{
			BuildTreeWithIntArr([]int{1, 2}), BuildTreeWithIntArrNillable([]*int{util.IntPtr(1), nil, util.IntPtr(2)}),
		}, false},
		{"test4", args{
			BuildTreeWithIntArr([]int{1, 2}), BuildTreeWithQueueNillable([]*int{util.IntPtr(1), nil, util.IntPtr(2)}),
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
