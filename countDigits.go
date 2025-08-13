package main

import (
	"fmt"
	"math"
)

// math.Abs returns the absolute value of a number
// math.Floor returns the largest integer less than or equal to a given number
// math.Round returns the nearest integer to a given number
// math.Ceil returns the smallest integer greater than or equal to a given number

func countDigits(n int) int {
	count := 0
	if n == 0 {
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

func CountDigitMain() {
	// Example usage
	// number := 0 // Output: 1
	// number := -123 // Output: 3
	number := 12245 // Output: 5
	result := countDigits(number)
	fmt.Printf("Number of Digits in number %v: %d \n", number, result)
}
