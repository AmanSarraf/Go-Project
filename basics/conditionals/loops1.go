package main

import "fmt"

func Loops() {
	fmt.Println("Enter a number")
	var num, i uint
	fmt.Scanln(&num)
	fmt.Printf("Numbers evenly divisible by %d between 1 and 100", num)

	for i = num; i <= 100; i++ {
		if i%num == 0 {
			fmt.Println(i)
		}
	}

}
