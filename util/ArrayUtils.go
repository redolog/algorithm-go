package util

// 不管顺序如何，只要a、b数组中元素均存在，并且出现次数一致，即判定两数组相等
func IntArrContainsEqual(a, b []int) bool {
	if a == nil {
		return b == nil
	}
	aLen, bLen := len(a), len(b)
	if aLen != bLen {
		return false
	}
	aNum2Cnt := make(map[int]int)
	bNum2Cnt := make(map[int]int)
	for i := 0; i < aLen; i++ {
		aNum2Cnt[a[i]]++
		bNum2Cnt[b[i]]++
	}

	for num, cnt := range aNum2Cnt {
		bCnt := bNum2Cnt[num]
		if cnt != bCnt {
			return false
		}
	}
	return true
}

func DeepEquals(a, b []float64, errorRange float64) bool {
	aLen := len(a)
	bLen := len(b)
	if aLen != bLen {
		return false
	}
	for i := 0; i < aLen; i++ {
		if !Equals(a[i], b[i], errorRange) {
			return false
		}
	}
	return true
}
