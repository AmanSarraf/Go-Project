package main

import "fmt"

func maps() {
	fmt.Print("This exapmple will show use of MAP")
	weekdays := make(map[string]string)
	weekdays["mon"] = "Monday"
	weekdays["tue"] = "Tuesday"
	weekdays["wed"] = "Wednesday"
	weekdays["thu"] = "Thrusday"
	weekdays["fri"] = "Friday"
	weekdays["sat"] = "Saturday"
	weekdays["sun"] = "Sunday"

	fmt.Println(weekdays)

}
