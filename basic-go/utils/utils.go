package utils

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strings"
)

// Handle_Error handles error with erroMessage
// and error
func Handle_Error(ErrMessage string, err error) {
	if err != nil {
		log.Fatal(ErrMessage, " "+err.Error())
	}
}

// GetInput returns cli data
func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(data), nil

}

// CheckFile takes fileinput and cheks if it
// is avaliable or not
func CheckFile(file string) fs.FileInfo {

	fileName, err := os.Stat(file)
	if err != nil {
		log.Fatal("file not found", err)
	}
	return fileName
}
