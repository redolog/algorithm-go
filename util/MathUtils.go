package util

import "math"

// 输入int数值，返回数值对应指针
func IntPtr(i int) *int {
	return &i
}

// 输入对象类型，转换为 []*int 数组形式
func Obj2IntegerBatch(objList []any) []*int {
	integerArr := []*int{}
	for _, obj := range objList {
		var num int
		if obj == nil {
			integerArr = append(integerArr, nil)
		} else {
			num = obj.(int)
			integerArr = append(integerArr, &num)
		}
	}
	return integerArr
}

func Equals(f1, f2, errorRange float64) bool {
	diff := math.Abs(f1 - f2)
	return diff < errorRange
}
