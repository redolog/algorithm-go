package util

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
