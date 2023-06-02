package util

// 输入int数值，返回数值对应指针
func IntPtr(i int) *int {
	return &i
}

// 输入对象类型，转换为 []*int 数组形式
func Obj2IntegerBatch(objList []any) []*int {
	integerArr := []*int{}
	for _, obj := range objList {
		if obj == nil {
			integerArr = append(integerArr, nil)
		} else {
			integerArr = append(integerArr, obj.(*int))
		}
	}
	return integerArr
}
