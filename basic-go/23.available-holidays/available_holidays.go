package main

import (
	"basic-go/date"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	holidayMap := make(map[time.Time]bool)
	countWeekends := 0

	fmt.Println("enter details of holidays in the following format", date.DATE_MONTH)
	
	for {
		data, err := getInput()
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
	//looping till startDate becomes endDate
	for endDate.After(startDate) {
		day := strconv.Itoa(startDate.Day())
		if len(day) == 1 {
			day = "0" + day
		}
		//parsing date and month only check with hashmap Data
		chekDate, err := time.Parse(string(date.DATE_MONTH_FULL), day+" "+startDate.Month().String())
		if err != nil {
			log.Fatal("error parsing date", err)
			break
		}
		//checking the given date is a weekend or holiday 
		if isWeekend(chekDate, holidayMap) {
			countWeekends++
		}
		startDate = startDate.AddDate(0, 0, 1)

	}
	fmt.Println("Available weekneds including holidays are", countWeekends)

}

func isWeekend(checkDate time.Time, holidayList map[time.Time]bool) bool {

	if checkDate.Weekday() == time.Saturday ||
		checkDate.Weekday() == time.Sunday || holidayList[checkDate] {
		return true
	}
	return false

}
func getInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(data), nil

}
