package main

import "fmt"

func array() {
	fmt.Println("This exampel will be Carrying array")

	fruits := []string{"apple", "mango", "banana", "guava"}
	fruits = append(fruits, "litchi")

	for index, value := range fruits {

		fmt.Printf("for Index = %v Fruit = %v\n", index, value)

	}

}
