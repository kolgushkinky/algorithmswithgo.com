package module01

import "fmt"

// FizzBuzz prints out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "FizzBuzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.

func FizzBuzz(n int) {

	for i := 1; i <= n; i++ {

		var isFizz = (i%3 == 0)
		var isBuzz = (i%5 == 0)
		var isFizzBuzz = (isFizz && isBuzz)

		if isFizzBuzz {
			fmt.Print("Fizz Buzz")
		} else if isFizz {
			fmt.Print("Fizz")
		} else if isBuzz {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Print("\n")
}
