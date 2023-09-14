package main

import (
	"basic-go/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("input list data")

	list := inputTillDone()
	fmt.Println("enter the name to lookup")
	searchData := make([]string, 0)
	pattern, err := utils.GetInput()
	if err != nil {	
		log.Fatal("error reading data")
	}

	for i := range list {
		//ranging listMap and appending to slice if matches
		if regexp.MustCompile(pattern).MatchString(i) {
			searchData = append(searchData, i)
		}
	}
	fmt.Println(searchData)
}
//inputting data to hashmap till done input occures
func inputTillDone() map[string]bool {
	listMap := make(map[string]bool)
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		if data == "done" || err != nil {
		
			break
		}
		listMap[strings.ToLower(data)] = true
	}
	return listMap
}
