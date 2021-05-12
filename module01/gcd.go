package module01

var (
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
)

func GCD(a, b int) int {
	aa := Factor(primes, a)
	bb := Factor(primes, b)

	m := make(map[int]int)

	for _, n := range bb {
		m[n] += 1
	}

	res := 1

	for _, n := range aa {
		if t, found := m[n]; found {
			if t > 0 {
				res *= n
				m[n] -= 1
			}
		}
	}

	return res
}
