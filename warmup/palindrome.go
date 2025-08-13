package warmup

func isPalindrome(num int) bool {
	original := num
	reversed := 0

	if num < 0 {
		return false // Negative numbers are not palindromes
	}

	for num > 0 {
		remainder := num % 10              // Get the last digit
		reversed = reversed*10 + remainder // Append the last digit to the reversed number
		num /= 10                          // Remove the last digit
	}

	return original == reversed
}

func PalindromeTest() {
	testCases := []int{121, -121, 10, 12321, 0, 1234321}
	resultsExpected := []bool{true, false, false, true, true, true}

	println("Palindrome Test Cases:")
	for i, n := range testCases {
		result := isPalindrome(n)
		if result == resultsExpected[i] {
			println("Test case", i+1, "(", n, "): Passed (Expected:", resultsExpected[i], ", Got:", result, ")")
		} else {
			println("Test case", i+1, "(", n, "): Failed (Expected:", resultsExpected[i], ", Got:", result, ")")
		}
	}
}
