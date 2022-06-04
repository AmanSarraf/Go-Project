package main

import "fmt"

func main() {

	fmt.Println("Hello and Welcome ! let's see different types of variables and datatypes")

	var name string = "Aman"
	surname := "sarraf"
	var mobile int64

	mobile = 9934419763
	var pi float64
	pi = 3.14024578

	fmt.Printf("Name is %s and surname is %s \n mobile number of %s is %d and value of pi is %.2f ", name, surname, name+surname, mobile, pi)

}
