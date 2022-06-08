package main

import (
	"bufio"
	"fmt"
	"os"
)

func userinput() {

	welcome := "Welcome to the the coffee"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please rate us")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	fmt.Printf("The type of input is %T", input)

}
