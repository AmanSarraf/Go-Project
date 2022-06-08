package main

import (
	"fmt"
	"strings"
)

type Person struct { //struct definition
	Firstname string
	Lastname  string
}

func Ugetter(P *Person) { //fucntion using struct

	P.Firstname = strings.ToUpper(P.Firstname)
	P.Lastname = strings.ToUpper(P.Lastname)

}

func main() {
	//struct as value
	var per1 Person
	per1.Firstname = "aman"
	per1.Lastname = "sarraf"
	Ugetter(&per1)
	fmt.Printf("Firstname = %s\nLastname = %s\n", per1.Firstname, per1.Lastname)

	//struct as pointer

	var per2 *Person = new(Person)
	per2.Firstname = "Pablo"
	per2.Lastname = "Nice Guy"
	Ugetter(per2)
	fmt.Printf("Firstname = %s\nLastname = %s\n", per2.Firstname, per2.Lastname)

	//struct as literals
	per3 := &Person{"sebestian", "cool boss"}
	Ugetter(per3)
	fmt.Printf("Firstname = %s\nLastname = %s\n", per3.Firstname, per3.Lastname)

}
