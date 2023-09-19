package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("input list data")
	list := inputTillDone()
	fmt.Println("enter the name to lookup")
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	if list[strings.ToLower(data)] {
		fmt.Println("name exists")
	} else {
		fmt.Println("name not exists")
	}

}

// inputTillDone inputs data from cli till
// done command given
func inputTillDone() map[string]bool {
	listMap := make(map[string]bool)
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		if data == "done" || err != nil {
			fmt.Println("here")
			break
		}
		listMap[strings.ToLower(data)] = true
	}
	return listMap
}
