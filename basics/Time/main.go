package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Current time is ")

	PresentTime := time.Now()
	fmt.Println(PresentTime.Format("01-02-2006 Monday"))
	fmt.Printf("formate is %T", PresentTime)

}
