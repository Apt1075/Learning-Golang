package main

import "fmt"

func main() {
	fmt.Println("Welcome to slice in GoLang")
	Numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Numbers)
	fmt.Println("Length is : ", len(Numbers))
	fmt.Println("Capacity is : ", cap(Numbers))

	EvenNumbers := []int{2, 4, 6, 8, 10}
	OddNumbers := []int{1, 3, 5, 7, 9, 11, 13}
	EvenNumbers = append(EvenNumbers, OddNumbers...)
	fmt.Println("Even Numbers : ", EvenNumbers)
	fmt.Println("Length is : ", len(EvenNumbers))
	fmt.Println("Capacity is : ", cap(EvenNumbers))

	Cars := make([]string, 4, 10)
	Cars = append(Cars, "Audi", "BMW", "Mercedes", "Mustang")
	fmt.Println(Cars)
	fmt.Println("Length is : ", len(Cars))
	fmt.Println("Capacity is : ", cap(Cars))

}
