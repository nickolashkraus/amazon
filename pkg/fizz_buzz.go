package pkg

import "fmt"

/*
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func fizzBuzz(n int32) {
	// Constraints:
	//   * 0 < n < 2 * 10^5
	//
	// If n does not satisfy the constraint, return.
	if n <= 0 || n >= 200000 {
		return
	}

	for i := int32(1); i <= n; i++ {
		if i%3 == 0 && i%5 == 0 { // Can also use i%15
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}

func FizzBuzz(n int32) {
	fizzBuzz(n)
}
