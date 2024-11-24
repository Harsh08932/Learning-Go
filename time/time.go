package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println("current time is ", currTime)
	fmt.Printf("The datatype for current time is %T\n", currTime)

	currDate := currTime.Format("2006/02/01 ,3:04 PM,Monday")
	fmt.Println("current Date is ", currDate)

	datestr := "2023-11-05,12:50 AM"
	layoutstr := "2006-02-01,3:04 PM"

	formattedTime, _ := time.Parse(layoutstr, datestr)
	fmt.Println("current time is ", formattedTime)

	newDate := currTime.Add(24 * time.Minute)
	fmt.Println(newDate)

}
