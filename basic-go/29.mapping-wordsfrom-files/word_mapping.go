package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	cmdArgs := os.Args
	//reading arguments including location and filename 
	fileInfo, err := os.Stat(strings.Join(cmdArgs[1:3], ""))
	if err != nil {
		log.Fatal("file not found ")
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(readAndPut(fileInfo.Name()))

}

// readAndPut reads fileData and stores into a map based on the repeatation
func readAndPut(fileData string) (map[string]int, error) {
	wordMap := make(map[string]int)
	file, err := os.Open(fileData)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if _, ok := wordMap[strings.ToLower(scanner.Text())]; ok {
			wordMap[strings.ToLower(scanner.Text())]++
		} else {
			wordMap[strings.ToLower(scanner.Text())] = 1
		}

	}
	return wordMap, nil
}
