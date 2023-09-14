package main

import (
	"basic-go/utils"
	"compress/gzip"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	OUTPUT_HTML_FILE = "merce-homepage.html"
)

func main() {
	//if no url provided panic occures recover panic situation
	defer func() {

		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				log.Println("no valid url provided")
			}
		}

	}()
	urlArgs := os.Args
	//checks url iis valid
	url, err := url.ParseRequestURI(urlArgs[1])
	if err != nil {
		log.Fatal("not a valid url")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//http client framework to manage http requests
	client := resty.New()

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	_, err = client.R().SetContext(ctx).SetOutput(OUTPUT_HTML_FILE).Get(url.String())
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatal("request timed out")

		}
		log.Fatal("error fetching details", err)
	}

	fmt.Println("fetched and stored webpage data")
	//reading and converting to compressed format
	zipFile, err := ReadAndCompress(OUTPUT_HTML_FILE)
	if err != nil {
		log.Fatal("cant complete operation", err)
	}
	fmt.Println("completed Read and Compression operations")
	fileHtml := utils.CheckFile(OUTPUT_HTML_FILE)
	fmt.Println("Before Compression ", fileHtml.Size())
	fileZip := utils.CheckFile(zipFile)
	fmt.Println("After compression", fileZip.Size())

}

// ReadAndCompress compress given file and returns compressed location
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
