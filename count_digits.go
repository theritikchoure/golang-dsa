package main

import "fmt"

func countDigits(nums int) int {
	count := 0

	for nums > 0 {
		count++
		nums = nums / 10
	}

	return count
}

func main() {
	nums := 77894
	counts := countDigits(nums)
	fmt.Println(counts)
}
