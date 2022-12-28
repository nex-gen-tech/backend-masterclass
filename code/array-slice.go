package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// Arrays
	a := []int{1, 4}
	a1 := a[0]
	a2 := a[1]
	a[0] = 6
	a[1] = 7 // way of adding a value to an slice or array

	a = append(a, 8)  // way of adding a value to an slice not in array
	a = append(a, 9)  // way of adding a value to an slice not in array
	a = append(a, 10) // way of adding a value to an slice not in array
	a = append(a, 11) // way of adding a value to an slice not in array

	fmt.Printf("type %T value %v \n", a, a)
	fmt.Printf("type %T value %v \n", a1, a1)
	fmt.Printf("type %T value %v \n", a2, a2)
	// fmt.Printf("ba - type %T value %v \n", ba, ba)

	// aCopy := make([]int, len(a))
	// copy(aCopy, a)

	i := a[:2]

	// loop through array
	for index, value := range a {
		fmt.Println(index, value)
	}

	fmt.Println(i)
}
