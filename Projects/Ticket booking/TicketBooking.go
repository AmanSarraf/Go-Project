package main

import (
	"fmt"
	"strings"
)

func main() {

	ConferenceName := "Go conference"
	// availableticket:= 50
	remainingticket := 50

	var Ticketbook uint

	fmt.Println("Welcome to ", ConferenceName, " Ticket booking booking app ", remainingticket, " are remaining")

	var Firstname string
	var Lastname string
	var bookings []string

	for remainingticket > int(Ticketbook) {
		fmt.Println("Enter your FirstName")
		fmt.Scanln(&Firstname)
		fmt.Println("Enter your LastName")
		fmt.Scanln(&Lastname)
		bookings = append(bookings, Firstname+" "+Lastname)
		fmt.Println("No. of tickets you want to book")
		fmt.Scanln(&Ticketbook)

		if Ticketbook > uint(remainingticket) {
			fmt.Println("can't book insufficient tickets remaining")

		} else {

			remainingticket -= int(Ticketbook)
			firstname := []string{}
			// var booking string

			// for i := 0; i < len(bookings); i++ { // this logic is used to iterate through each of a slice
			// 	booking = bookings[i]
			// 	names := strings.Split(booking, " ")
			// 	firstname = append(firstname, names[0])
			// }

			for _, first := range bookings {

				// String s is split on the basis of white spaces
				// and store in a string array
				var name = strings.Fields(first)
				var booking = name[0]
				firstname = append(firstname, booking)

			}

			fmt.Printf("Booking Succcessful for %s  and no of tickets booked is =%d, remaning Tickets=%d", firstname, Ticketbook, remainingticket)
		}

	}
}
