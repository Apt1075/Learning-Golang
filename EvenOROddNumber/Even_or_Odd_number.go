package main

import "fmt"

func main() {
	fmt.Println("Enter The Number")
	var num int
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Number is Even")
	} else if num == 0 {
		fmt.Println("Number is Zero")
	} else {
		fmt.Println("Number is Odd")
	}
}
