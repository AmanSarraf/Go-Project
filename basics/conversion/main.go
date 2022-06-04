package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	welcome := "Welcome to our Pizza shop"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate on scale of 1 to 5")
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)
	fmt.Printf("The type of rating is %T", input)

	numeric_rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	//
	if err != nil {

		fmt.Print(err)
	} else {
		fmt.Println(numeric_rating + 1)
	}

}
