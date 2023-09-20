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
	fmt.Println("enter a date in the format of ", date.DD_MM_YY)
	inputDate, err := date.GetInputDate(date.DD_MM_YY)
	if err != nil {
		log.Fatal("error parsing date")
	}
	days, err := utils.GetInput()
	if err != nil {
		log.Fatal("error reading data")
	}
	buissinessDay, err := strconv.Atoi(days)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CheckDates(inputDate, buissinessDay, holidayMap))
}
//CheckDates returns the relative day after comparing given inputDate and business days
func CheckDates(inputDate time.Time, buissinessDays int, holidayMap map[time.Time]bool) (string, error) {
	if inputDate.IsZero() {
		return "", errors.New("provide valid time")
	}

	if buissinessDays > 0 {
		for i := buissinessDays; i > 0; {
			inputDate = inputDate.AddDate(0, 0, 1)
			if buissinessDays%5 == 0 {
				inputDate = inputDate.AddDate(0, 0, 7)
				i -= 5
				continue
			}
			if isHoliday(inputDate, holidayMap) {
				continue
			}
			i--

		}

	} else if buissinessDays < 0 {
		for i := buissinessDays; i < 0; {
			inputDate = inputDate.AddDate(0, 0, -1)
			if buissinessDays%5 == 0 {
				inputDate = inputDate.AddDate(0, 0, 7)
				i -= 5
				continue
			}
			if isHoliday(inputDate, holidayMap) {
				continue
			}
			i++
		}
	}
	return inputDate.Format(string(date.DD_MM_YY)), nil
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
