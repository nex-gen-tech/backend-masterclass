package code

import "fmt"

// Write a function that takes an arbitrary number of ints and returns their sum.
func sum(nums ...int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	val := sum(numbers...) // 1,2,3,4,5,6,7,8,9,10
	fmt.Println(val)
}
