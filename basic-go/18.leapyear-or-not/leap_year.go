package main

import (
	"basic-go/date"
	"basic-go/utils"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("enter the date in the following format ", date.DD_MM_YY)
	data, err := utils.GetInput()
	if err != nil {
		log.Fatal("error reading data", err)

	}

	time, err := time.Parse(string(date.DD_MM_YY), data)
	if err != nil {
		log.Fatal("error parsing date time")
	}

	if !LeapYearOrNot(time) {
		fmt.Println("not a leap year")
		return
	}
	fmt.Println("its a leap year")

}
//LeapYearOrNot checks given date is leapYear or not
func LeapYearOrNot(time time.Time) bool {
//leap year checking logic
	if time.Year()%4 == 0 || time.Year()%100 != 0 {
		return true
	}
	return false
}
