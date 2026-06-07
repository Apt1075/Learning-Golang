package main

import "fmt"

func Helloworld() {
	fmt.Println("Helloworld")
}

func AddtionTwoNumber() {
	var num1 int = 20
	var num2 int = 30
	var sum int
	sum = num1 + num2
	println("Sum of Two Numbers == ", sum)
}

func ReminderTwoNumber(a, b int) {
	if b == 0 {
		fmt.Println("Number is Zero")
		return
	}
	var Reminder = a % b
	println("Reminder of Two Numbers", Reminder)
	// return Reminder
}

func DivisionTwoNumber(a, b float64) int {
	if b == 0 {
		fmt.Println("Number is Zero")
		return 0
	}
	Division := a / b
	// println("Division of Two Numbers", Division)
	return int(Division)
}

func main() {
	fmt.Println("We are learning function in GoLang")
	Helloworld()
	AddtionTwoNumber()
	ReminderTwoNumber(10, 3)
	Divide := DivisionTwoNumber(10, 3)
	println("Division of Two Numbers", Divide)
}
