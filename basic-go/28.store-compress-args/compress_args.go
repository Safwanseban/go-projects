package main

import (
	"basic-go/utils"
	"compress/gzip"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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
	//checks url is valid
	url, err := url.ParseRequestURI(urlArgs[1])
	if err != nil {
		log.Fatal("not a valid url")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := FetchAndSetOutput(ctx, url.String(), OUTPUT_HTML_FILE); err != nil {
		if ctx.Err() != context.DeadlineExceeded {
			log.Fatal("error fetching details")
		}
		log.Fatal(err)
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

func FetchAndSetOutput(ctx context.Context, url string, outPutFile string) (int, error) {
	client := resty.New().SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	resp, err := client.R().SetContext(ctx).SetOutput(outPutFile).Get(url)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return http.StatusRequestTimeout, errors.New("request timed out")
		}
		return resp.StatusCode(), err
	}
	return resp.StatusCode(), nil

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
