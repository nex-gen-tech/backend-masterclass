package main

import "fmt"

func main() {
	result := frequency([]string{"a", "b", "c", "a", "b"})
	fmt.Println(result)
}

// 1. Write a function that takes a slice of strings and returns a map with the frequency of each string in the slice.
// For example, given the ->
// input []string{"a", "b", "c", "a", "b"},
// the function should return a map { "a": 2, "b": 2, "c": 1 }.

// Input -> []string{"a", "b", "c", "a", "b"}
// Output -> map[string]int{"a": 2, "b": 2, "c": 1}
func frequency(s []string) map[string]int {
	// define a map
	res := make(map[string]int)
	// res := map[string]int{}
	// var res map[string]int

	// Iterate over the slice of strings
	for _, v := range s {
		// fmt.Println(i, v)
		// Check if the string is already in the map
		val, ok := res[v]
		if ok {
			// If it is, increment the value
			res[v] = val + 1
		} else {
			// If it is not, add it to the map with a value of 1
			res[v] = 1
		}
	}

	return res
}
