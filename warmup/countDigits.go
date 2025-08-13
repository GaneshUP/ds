package warmup

import (
	"fmt"
	"math"
)

// math.Abs returns the absolute value of a number
// math.Floor returns the largest integer less than or equal to a given number
// math.Round returns the nearest integer to a given number
// math.Ceil returns the smallest integer greater than or equal to a given number

// countDigits counts the number of digits in an integer n
func countDigits(n int) int {
	count := 0
	if n == 0 { // Special case for 0
		return 1
	}
	n = int(math.Abs(float64(n))) // Ensure n is positive
	// Count the number of digits
	for n > 0 {
		n = int(math.Floor(float64(n) / 10))
		count++
	}
	return count
}

func CountDigitTest() {
	testCases := []int{0, -1, 123456789, 1000000, -9999, 42, 7, -1234567890}
	resultExpected := []int{1, 1, 9, 7, 4, 2, 1, 10}
	fmt.Println("Count Digit Test Cases:")
	// Loop through test cases and print results
	for i, n := range testCases {
		count := countDigits(n)
		if count == resultExpected[i] {
			fmt.Printf("Test case %d (%d): Passed (Expected: %d, Got: %d)\n", i+1, n, resultExpected[i], count)
		} else {
			fmt.Printf("Test case %d (%d): Failed (Expected: %d, Got: %d)\n", i+1, n, resultExpected[i], count)
		}

	}
}
