package decimal

// https://leetcode.cn/problems/convert-the-temperature/
// 2469. Convert the Temperature

// Kelvin = Celsius + 273.15
// Fahrenheit = Celsius * 1.80 + 32.00
func convertTemperature(celsius float64) []float64 {
	return []float64{
		celsius + 273.15,
		celsius*1.80 + 32.00,
	}
}
