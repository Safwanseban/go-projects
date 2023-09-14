package main

import (
	"basic-go/utils"
	"compress/gzip"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	URL_MERCE         = "https://merce.co"
	OUTPUT_HTML_FILE  = "merce-homepage.html"
	DEFAULT_DIRECTORY = "/home/merceadm/safwan/assignments/basic-go/27.storeandcompress/downloads"
)

func main() {

	client := resty.New()
	client.SetOutputDirectory(DEFAULT_DIRECTORY)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	oFileLoc := fmt.Sprintf("%s/%s", DEFAULT_DIRECTORY, OUTPUT_HTML_FILE)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := client.R().SetContext(ctx).SetOutput(oFileLoc).Get(URL_MERCE)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatal("request timed out")
		}
		log.Fatal("error fetching details", err)
	}
	fmt.Println("fetched and stored webpage data")

	zipFile, err := ReadAndCompress(oFileLoc)
	if err != nil {
		log.Fatal("cant complete operation", err)
	}

	fmt.Println("completed Read and Compression operations")
	fileHTml := utils.CheckFile(oFileLoc)
	fmt.Println("Before Compression ", fileHTml.Size())

	fileZip := utils.CheckFile(zipFile)
	fmt.Println("After compression", fileZip.Size())

}

func ReadAndCompress(file string) (string, error) {
	newFile, err := os.Open(file)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(newFile)
	if err != nil {
		return "", err
	}
	file_name := file + ".gz"
	outputFile, err := os.Create(file_name)
	if err != nil {
		return "", err
	}
	w := gzip.NewWriter(outputFile)
	w.Write(data)
	w.Close()
	return outputFile.Name(), nil
}
