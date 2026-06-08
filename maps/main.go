package main

import "fmt"

func main() {
	StudentName := make(map[int]string)
	StudentName[1] = "Arpit"
	StudentName[2] = "Kajal"
	StudentName[3] = "Pulkit"
	fmt.Println(StudentName)

	value, key_exits := StudentName[2]
	if key_exits {
		fmt.Println("Key exists")
		fmt.Println("value == ", value)
	} else {
		fmt.Println("Key does not exist")
	}

	// EmployeeAge := make(map[string]int)
	// EmployeeAge["Arpit"] = 21
	// EmployeeAge["Kajal"] = 22
	// EmployeeAge["Pulkit"] = 23
	// fmt.Println(EmployeeAge)

	// for key, value := range StudentName {
	// 	fmt.Println(key, value)
	// }

	// for key, value := range EmployeeAge {
	// 	fmt.Println(key, value)
	// }
}
