package main

import "fmt"

func main() {
	fmt.Println("If else in GoLang")
	if 5%2 == 0 {
		fmt.Println("The value of 5 is Even")
	} else {
		fmt.Println("The value of 5 is Odd")
	}

	if num := 10; num%2 == 0 {
		fmt.Println("The value of 10 is Even")
	} else if num := 10; num%2 != 0 {
		fmt.Println("The value of 10 is Odd")
	} else if num == 0 {
		fmt.Println("The value of 0 is Zero")
	}

}
