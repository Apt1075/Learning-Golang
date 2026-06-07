package main

import "fmt"

func DivisionTwoNumber(a, b float64) (float64, error) {
	if b == 0 {
		fmt.Println("Number is Zero")
		return 0, fmt.Errorf("demonstartion must not be 0")
	}
	Division := a / b
	// println("Division of Two Numbers", Division)
	return Division, nil
}

// func main() {
// 	fmt.Println("Welcome to Error Handling in GoLang")
// 	Divide, err := DivisionTwoNumber(10, 3)
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	println("Division of Two Numbers", Divide)
// }
func main() {
	fmt.Println("Welcome to Error Handling in GoLang")
	Divide, _ := DivisionTwoNumber(10, 3)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	println("Division of Two Numbers", Divide)
}
