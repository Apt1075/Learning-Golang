package main

import "fmt"

func ForLoop() {

	// 1. Normal for loop
	fmt.Println("1. Normal for loop")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n2. For loop without initializer and post statement (while style)")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n3. For loop without any condition (infinite loop with break)")
	for j := 1; ; j++ {
		if j > 5 {
			break
		}
		fmt.Println(j)
	}

	fmt.Println("\n4. For loop as for-each (range over array)")
	numbers := [5]int{10, 20, 30, 40, 50}
	for _, num := range numbers {
		fmt.Println(num)
	}

	fmt.Println("\n5. For loop as for-each (range over slice)")
	slices := []int{60, 70, 80, 90, 100}

	for _, slice := range slices {
		fmt.Println(slice)
	}

	fmt.Println("\n6. Nested for loop")
	for k := 1; k <= 3; k++ {
		for l := 1; l <= k; l++ {
			fmt.Printf("%d ", l)
		}
		fmt.Println()
	}

	fmt.Println("\n7. For loop with continue")
	for m := 1; m <= 5; m++ {
		if m == 3 {
			continue
		}
		fmt.Println(m)
	}
}

func main() {
	fmt.Println("For loop is learning in Golang")
	ForLoop()
}
