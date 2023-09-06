package main

import (
	"fmt"
	"math"
)

func checkArmstrong(x int) bool {
	var originalNum int = x
	var sum int = 0

	for x > 0 {
		lastDigit := x % 10

		sum = sum + int(math.Pow(float64(lastDigit), 3))

		x = x / 10
	}

	return sum == originalNum
}

const armstrongNums int = 371

func main() {

	var isArmstrong bool = checkArmstrong(armstrongNums)

	fmt.Println(isArmstrong)
}
