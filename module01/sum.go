package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {

	if numbers == nil {
		return 0
	}

	sum := 0

	for _, x := range numbers {
		sum += x
	}

	return sum
}
