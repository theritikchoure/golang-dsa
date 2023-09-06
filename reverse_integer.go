// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:
// Input: x = 123
// Output: 321

// Example 2:
// Input: x = -123
// Output: -321

// Example 3:
// Input: x = 120
// Output: 21

// Constraints: -231 <= x <= 231 - 1

// ----------------------------------------------------------------------------------------------------------------

package main

import "fmt"

func reverse(nums int) int {

	// Constants for integer limits
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)

	// Reverse integer variable
	var reverseInt int

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
			return 0
		}

		if reverseInt < MinInt32/10 || (reverseInt == MinInt32/10 && lastDigit < -8) {
			return 0
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

	// return reverseInt
	return reverseInt
}

func main() {
	nums := 1534236469
	counts := reverse(nums)
	fmt.Println(counts)
}
