package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_old() {
	// fmt.Println("Hello Word")
	// test.Helloworld()

	// var Name string = "Arpit"
	// println("name - ", Name)

	// var Age int = 28
	// println("age= ", Age)

	// var IsMarried bool = false
	// println("IsMarried  = ", IsMarried)

	// var Salary float64 = 45000.00
	// println("Salary = ", Salary)

	// const Address string = "Ahmedabad"
	// println("Address = ", Address)

	// Student := false
	// println("Student = ", Student)
	// fmt.Println("Name=", Name, "Age=", Age, "Salary=", Salary)

	// fmt.Printf("Name =%s, Age=%d, Salary=%.3f", Name, Age, Salary)
	var Name string
	// fmt.Scanf("%s\n", &Name)
	// fmt.Println("Name=", Name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name and Age")
	Name, _ = reader.ReadString('\n')
	fmt.Println("Name=", Name)

}
