package main

import (
	"fmt"
	"math"
)

func checkPrime(nums int) bool {
	var count int

	for i := 1; i <= int(math.Sqrt(float64(nums))); i++ {
		if nums%i == 0 {
			count++

			if nums/i != i {
				count++
			}
		}
	}

	return count == 2
}

func main() {
	nums := 11
	fmt.Println(checkPrime(nums))
}
