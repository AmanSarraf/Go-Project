package main

import (
	"fmt"
	"strings"
)

func slice1() {
	fmt.Println("hello")
	sample := []string{"aman kumar", "raman baba", "sunil mahto", "mohan prasad"}
	var names []string
	for _, first := range sample {

		// String s is split on the basis of white spaces
		// and store in a string array
		var x = strings.Fields(first)
		var n = x[0]
		names = append(names, n)

	}
	fmt.Printf("First name = %s ", names)

}
