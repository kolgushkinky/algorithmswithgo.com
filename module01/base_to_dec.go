package module01

import "math"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	//	var digits string = "0123456789ABCDEF"

	var res int
	l := len(value)

	for i, r := range value {
		var digit int
		if r < 65 {
			digit = int(r - '0')
		} else {
			digit = int(r - 55)
		}
		res += digit * int(math.Pow(float64(base), float64(l-i-1)))
	}
	return res
}
