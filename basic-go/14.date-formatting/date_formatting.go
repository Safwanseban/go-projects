// Write a program to print current date/time in following formats
//
//	(one line per format) in UTC time
//	 e.g. “16 Mar 2022” “Mar 16, 2022” “2022-03-16” “2022-03-16T15:52:00Z” “Tuesday,
//	  16 March 2022” [Date manipulation, 3 hours]
package main

import (
	"fmt"
	"time"
)

type date string

const (
	DATE_MONTH_YEAR       date = "02 Jan 2006"
	MONTH_DATE_YEAR       date = "Jan 02 2006"
	YY_MM_DD              date = "2006-01-02"
	ZONE_TIME_FORMAT      date = "2006-01-02T15:04:05Z"
	DATE_NDATE_MONTH_YEAR date = "Monday, 02 January 2006"
)

func main() {
	fmt.Println(getFormattedTime(string(DATE_MONTH_YEAR)))
	fmt.Println(getFormattedTime(string(MONTH_DATE_YEAR)))
	fmt.Println(getFormattedTime(string(YY_MM_DD)))
	fmt.Println(getFormattedTime(string(ZONE_TIME_FORMAT)))
	fmt.Println(getFormattedTime(string(DATE_NDATE_MONTH_YEAR)))
}

func getFormattedTime(layout string) string {

	return time.Now().Format(layout)

}
