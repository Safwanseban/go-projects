package main

import (
	"basic-go/date"
	"fmt"
	"log"
)

func main() {

	fmt.Println("Please enter date in the following format ", date.DD_MM_YY)
	firstDate, err := date.GetInputDate(date.DD_MM_YY)
	if err != nil {

		log.Fatal("error fetching date")
	}
	secondDate, err := date.GetInputDate(date.DD_MM_YY)
	if err != nil {
		log.Fatal("error fetching date")
	}

	duration := date.DateDiffrence(firstDate, secondDate)
	fmt.Println("Diffrence between two dates is ", duration)

}
