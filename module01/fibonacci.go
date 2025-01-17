package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func Fibonacci(n int) int {
	// switch {
	// case n == 0:
	// 	return 0
	// case n == 1:
	// 	return 1
	// default:
	// 	return Fibonacci(n-1) + Fibonacci(n-2)
	// }

	if n <= 1 {
		return n
	}

	n0 := 0
	n1 := 1

	var curVal int
	for i := 2; i <= n; i++ {
		curVal = n0 + n1
		n0 = n1
		n1 = curVal
	}

	return curVal
}
