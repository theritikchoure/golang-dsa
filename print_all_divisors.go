package main

import (
	"fmt"
	"math"
	"sort"
)

func printAllDivisors(nums int) {
	var numbers []int

	for i := 1; i <= int(math.Sqrt(float64(nums))); i++ {
		if nums%i == 0 {
			numbers = append(numbers, i)

			if nums/i != i {
				numbers = append(numbers, nums/i)
			}
		}
	}

	sort.Ints(numbers)

	fmt.Println(numbers)
}

func main() {
	nums := 27
	printAllDivisors(nums)
}
