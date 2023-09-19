package main

import (
	"basic-go/utils"
	"fmt"
	"log"
	"time"
)

var ValidTimeZones = map[string]bool{
	"Asia/Kolkata":     true,
	"Asia/Dubai":       true,
	"America/New_York": true,
	"Europe/London":    true,
}

func main() {
	fmt.Println("available time zones are")
	for zones := range ValidTimeZones {
		fmt.Println(zones)

	}
	fmt.Println("enter a location from the above")
	location, err := utils.GetInput()
	if err != nil {
		log.Fatal("error getting input")
	}
	if !ValidTimeZones[location] {
		log.Fatal("enter a timezone listed from the above")
	}
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Fatal("error loading location")
	}
	fmt.Println(time.Now().In(loc))

}
