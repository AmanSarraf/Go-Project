package main

import "fmt"

func userio() {

	var fname string
	var lname string
	var dob string
	var mobile int64
	fmt.Println("Please Enter your Details")
	fmt.Println("Firstname")
	fmt.Scanln(&fname)

	fmt.Println("Lastname")

	fmt.Scanln(&lname)

	fmt.Println("Dateofbirth")

	fmt.Scanln(&dob)
	fmt.Println("Mobile")

	fmt.Scanln(&mobile)

	fmt.Printf("Your Details are\nFullName=%s %s DOB=%s Mobile=%d", fname, lname, dob, mobile)
}
