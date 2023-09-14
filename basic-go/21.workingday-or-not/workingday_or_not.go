package main

import (
	"basic-go/date"
	"basic-go/utils"
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

	day := strconv.Itoa(checkDate.Day())
	if len(day) == 1 {
		day = "0" + day
	}
	checkDate, err = time.Parse(string(date.DATE_MONTH_FULL), day+" "+checkDate.Month().String())
	if err != nil {
		log.Fatal("error parsing date", err)

	}

	if !isWeekend(checkDate, holidayMap) {
		fmt.Println("entered date is a working day")
		return
	}
	fmt.Println("entered date is not a working day")

}
func isWeekend(checkDate time.Time, holidayList map[time.Time]bool) bool {

	if checkDate.Weekday() == time.Saturday ||
		checkDate.Weekday() == time.Sunday || holidayList[checkDate] {
		return true
	}
	return false

}
