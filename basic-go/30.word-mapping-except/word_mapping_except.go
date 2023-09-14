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
	fileInfo, err := os.Stat(strings.Join(cmdArgs[1:3], ""))
	if err != nil {
		log.Fatal("file not found ")
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(readAndPut(fileInfo.Name())) //reading and putting data to hashmap based on words

}
//readAndPut
func readAndPut(fileData string) (map[string]int, error) {
	wordMap := make(map[string]int)
	file, err := os.Open(fileData)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() { //scanning file till EOF occurs
		if CommonWords(scanner.Text()) { //scanning for common words if ignore
			continue
		}
		//if alreading in hashmap increment else add insert
		if _, ok := wordMap[strings.ToLower(scanner.Text())]; ok {
			wordMap[strings.ToLower(scanner.Text())]++
		} else {
			wordMap[strings.ToLower(scanner.Text())] = 1
		}

	}
	return wordMap, nil
}

// CommonWords checks if the words is available commonWords map
func CommonWords(s string) bool {

	commonWords := map[string]bool{

		"the": true,
		"a":   true,
		"an":  true,
		"i":   true,
	}
	return commonWords[strings.ToLower(s)]

}
