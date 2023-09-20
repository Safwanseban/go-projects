package main

import (
	"basic-go/date"
	"basic-go/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Assignment:21
// Program prompts for user to enter the holiday date in the format of 02 jan till done command.
// then prompts to check a date to check weather it is a working day or not.
// checks date is weekend or a date that is part of holiday list then returns the result

func main() {

	holidayMap := make(map[time.Time]bool)

	fmt.Println("enter details of holidays in the following format", date.DATE_MONTH)
	for {
		data, err := utils.GetInput()
		if err != nil || data == "done" {
			break
		} else {
			time, err := time.Parse(string(date.DATE_MONTH), data)
			if err != nil {
				log.Fatal("error in parsing time")
			}
			holidayMap[time] = true
		}
	}
	fmt.Printf("enter the date in this format %s to check working day or not\n", date.DD_MM_YY)
	checkDate, err := date.GetInputDate(date.DD_MM_YY)
	if err != nil {
		log.Fatal("error reading data")
	}
	//checks holiday or not
	if !isHoliday(checkDate, holidayMap) {
		fmt.Println("entered date is a working day")
		return
	}
	fmt.Println("entered date is not a working day")

}

// isHoliday checks given date is holiday or not
func isHoliday(checkDate time.Time, holidayList map[time.Time]bool) bool {
	if checkDate.Weekday() == time.Saturday ||
		checkDate.Weekday() == time.Sunday || holidayList[removeYear(checkDate)] {
		return true
	}
	return false

}

// removeYear removes year from a given date
func removeYear(checkDate time.Time) time.Time {
	var err error
	day := strconv.Itoa(checkDate.Day())
	if len(day) == 1 {
		day = "0" + day
	}
	checkDate, err = time.Parse(string(date.DATE_MONTH_FULL), day+" "+checkDate.Month().String())
	if err != nil {
		log.Fatal("error parsing date", err)

	}
	return checkDate
}
