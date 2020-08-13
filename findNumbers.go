package find_numbers

import (
	"math"
	"strconv"
)

func findNumbers(nums []int) int {
	result := 0
	for _, val := range nums {
		if calculateDigits(val)%2 == 0 {
			result += 1
		}
	}
	return result
}

func calculateDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func calculateDigitsV1(num int) int {
	result := strconv.Itoa(num)
	return len(result)
}
