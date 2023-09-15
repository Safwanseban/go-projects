package main

import (
	"basic-go/date"
	"basic-go/utils"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

/*
Assignment 23:
Program prompts for user to enter the holiday date in the format of 02 jan till done command.
Data is been appended to a holidayMap and new prompt for startDate and EndDate is given.
loops from startDate to EndDate, doing so checks the date is holiday or not and counts the holidayDays.
*/

func main() {
	holidayMap := make(map[time.Time]bool)

	fmt.Println("enter details of holidays in the following format", date.DATE_MONTH)

	for {
		data, err := utils.GetInput()
		if err != nil {
			break
		}
		if data == "done" {
			break
		} else {
			//parsing time and data to the given format
			time, err := time.Parse(string(date.DATE_MONTH), data)
			if err != nil {
				log.Fatal("error in parsing time")
			}
			holidayMap[time] = true
		}
	}

	fmt.Println("Please enter date in the following format ", date.DD_MM_YY)
	//getting dates to check
	startDate, err := date.GetInputDate(date.DD_MM_YY)
	if err != nil {
		log.Fatal("error fetching date data")
	}
	endDate, err := date.GetInputDate(date.DD_MM_YY)
	if err != nil {

		log.Fatal("error fetching date data")
	}
	holidays, err := CheckDates(startDate, endDate, holidayMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available holidays including weekends are", holidays)

}

// CheckDates two Date Arguments and HolidayList returns error if error is there
// else provides valid holidaycount
func CheckDates(startDate, endDate time.Time, holidayMap map[time.Time]bool) (int, error) {
	countWeekends := 0
	//checks if given time is valid or not
	if startDate.IsZero() || endDate.IsZero() {
		return 0, errors.New("please provide valid time")

	}
	if endDate.Before(startDate) {
		return 0, errors.New("endDate is before start date")
	}
	for endDate.After(startDate) {

		//checking the given date is a weekend or holiday
		if isHoliday(startDate, holidayMap) {
			countWeekends++
		}
		startDate = startDate.AddDate(0, 0, 1)

	}
	return countWeekends, nil

}

// isHoliday takes a date and holidaylist argument and returns the givenDate
// is holiday or not
func isHoliday(checkDate time.Time, holidayList map[time.Time]bool) bool {

	if checkDate.Weekday() == time.Saturday ||
		checkDate.Weekday() == time.Sunday || holidayList[removeYear(checkDate)] {
		return true
	}
	return false

}

// removeYear removes the given date year returns month and day only
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
