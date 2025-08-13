package warmup

import "math"

// reverseInteger reverses the digits of an integer n, limiting [-2^32, 2^32] range
func reverseInteger(n int) int {
	reversed := 0
	nCopy := n
	limit := int(math.Pow(2, 32)) // 2^32 is the limit for 32-bit signed integers

	n = int(math.Abs(float64(n))) // Ensure n is positive
	for n > 0 {
		reminder := n % 10                    // Get the last digit
		reversed = (reversed * 10) + reminder // Append the last digit to the reversed number
		n /= 10                               // Remove the last digit
	}

	if nCopy < 0 {
		reversed = -reversed // Restore the sign for negative numbers
	}

	if reversed < -limit || reversed > limit {
		return 0 // Return 0 if the reversed number is out of bounds
	}

	return reversed
}

func ReverseIntegerTest() {
	testCases := []int{123, -123, 120, 0, 1534236469, -2147483648}
	expectedResults := []int{321, -321, 21, 0, 0, 0}

	println("Reverse Integer Test Cases:")
	for i, n := range testCases {
		result := reverseInteger(n)
		if result == expectedResults[i] {
			println("Test case", i+1, "(", n, "): Passed (Expected:", expectedResults[i], ", Got:", result, ")")
		} else {
			println("Test case", i+1, "(", n, "): Failed (Expected:", expectedResults[i], ", Got:", result, ")")
		}
	}
}
