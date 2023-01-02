package code

import "fmt"

func main() {
	// need to store student name and age
	// let's say we have 3 students
	// rama, 20
	// ravi, 21
	// raju, 22

	// studentsName := []string{"rama", "ravi", "raju"}
	// studentsAge := []int{20, 21, 22}

	// for i := 0; i < len(studentsAge); i++ {
	// 	println(studentsName[i], studentsAge[i])
	// }

	// var dict map[key]value

	students := map[string]int{
		"rama": 20,
		"ravi": 21,
		"raju": 22,
	}

	// fmt.Println(students)

	// // Let's loop over the map
	// for key, value := range students {
	// 	fmt.Println(key, value)
	// }

	// Let's get the value of a key
	// fmt.Println(students["raju"])

	// Let's change Ravi's age
	// students["ravi"] = 25
	// students["rama"] = 31
	// fmt.Println(students)

	// Let's add new Student
	students["sita"] = 23
	students["anil"] = 44
	// fmt.Println(students)
	// fmt.Println(students["sita"])

	// Delete a Student
	// delete(students, "sita")
	// fmt.Println(students)
	// fmt.Println(students["sita"])

	// Let's check if a key exists
	value, ok := students["lata"]
	if ok {
		fmt.Println(value)
	}
}
