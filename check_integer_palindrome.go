package main

import "fmt"

func checkIntegerPalindrome(nums int) bool {

	// Constants for integer limits
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)

	// Reverse integer variable
	var reverseInt int
	var originalInt int = nums

	// Integer is negative or positive
	negative := false

	// Check for negative
	if nums < 0 {
		negative = true
		nums = -1 * nums // if given integer is negative then convert it to positive first
	}

	for nums > 0 {

		// Extract the last digit
		lastDigit := nums % 10

		// Check for integer overflow before updating reverseInt
		if reverseInt > MaxInt32/10 || (reverseInt == MaxInt32/10 && lastDigit > 7) {
			return false
		}

		if reverseInt < MinInt32/10 || (reverseInt == MinInt32/10 && lastDigit < -8) {
			return false
		}

		// Update result by shifting and adding the last digit
		reverseInt = (reverseInt * 10) + lastDigit

		// Remove the last digit from nums
		nums = nums / 10
	}

	// if given integer is negative then convert reverseInt to negative
	if negative {
		reverseInt = -1 * reverseInt
	}

	if reverseInt == originalInt {
		return true
	}

	return false
}

func main() {
	nums := 757
	counts := checkIntegerPalindrome(nums)
	fmt.Println(counts)
}
