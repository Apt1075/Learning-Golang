package main

import "fmt"

func positive_or_negative_number(n int) {
	if n > 0 {
		fmt.Println("Number is Positive")
	} else if n == 0 {
		fmt.Println("Number is Zero")
	} else {
		fmt.Println("Number is negative")
	}
}

func main() {
	fmt.Println("Enter the number")
	var num int
	fmt.Scan(&num)
	positive_or_negative_number(num)
}
