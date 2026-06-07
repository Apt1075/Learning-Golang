package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in GoLang")

	var cars = []string{"BMW", "TOYOTA", "AUDI"}
	fmt.Println(cars)
	fmt.Printf("Type of cars is %T", cars)

	var Number [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("\nNumber is :", Number)
}
