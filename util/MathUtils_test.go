package util

import (
	"reflect"
	"testing"
)

func TestObj2IntegerBatch(t *testing.T) {
	type args struct {
		objList []any
	}
	tests := []struct {
		name string
		args args
		want []*int
	}{
		{
			name: "test1",
			args: args{
				objList: []any{1, nil, 2},
			},
			want: []*int{IntPtr(1), nil, IntPtr(2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Obj2IntegerBatch(tt.args.objList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Obj2IntegerBatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
