package main

import "fmt"

func main() {

	fmt.Println("Enter the Number")
	var num int
	sum := 0
	fmt.Scan(&num)

	// for i := 1; i <= num; i++ {
	// 	sum += i
	// }
	sum = num * (num + 1) / 2
	fmt.Println(sum)
}
