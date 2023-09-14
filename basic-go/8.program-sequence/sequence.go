package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	END_SEQUENCE = "proceed"
)

func main() {

	var (
		input string
		sum int
	)
	for {
		fmt.Scanf("%s", &input)
		if input != END_SEQUENCE {
			data, err := strconv.Atoi(input)
			if err != nil {
				log.Fatal("not a valid number")
				break
			}
			sum += data
			
		} else {
			fmt.Println("result is", sum)
			break
		}

	}

}

