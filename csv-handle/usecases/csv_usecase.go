package usecases

import (
	client "csv-handle/client/interfaces"
	"csv-handle/usecases/interfaces"
	"csv-handle/utils"
	"encoding/csv"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"sync"
)

type CsvUsecases struct {
	Svc client.AwsClientI
}

func NewCsvUsecase(client client.AwsClientI) interfaces.CsvUsecase {
	return &CsvUsecases{
		Svc: client,
	}
}
func (cu *CsvUsecases) BatchUpload(files []*multipart.FileHeader) ([]map[string]any, int) {
	results := make([]map[string]any, 0)
	var wg sync.WaitGroup
	count := 0
	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader, wg *sync.WaitGroup) {
			defer wg.Done()
			oFile, err := file.Open()
			if err != nil {
				results = utils.ManageResult(results, file.Filename, "failed", errors.New("file opening error"))
				return
			}
			defer oFile.Close()
			fileName := strings.Split(file.Filename, ".")
			extension := fileName[len(fileName)-1]
			if extension != "csv" {
				count++
				results = utils.ManageResult(results, file.Filename, "failed", errors.New("not a valid extension"))
				return
			}
			if err := cu.Svc.UploadToS3("csvhandlebucket", utils.NewUUIDFile(extension), oFile); err != nil {
				results = utils.ManageResult(results, file.Filename, "failed", errors.New("can't upload to s3"))
				return
			}
			results = utils.ManageResult(results, file.Filename, "success", nil)

		}(file, &wg)
	}
	wg.Wait()
	return results, count

}

func (cu *CsvUsecases) ValidateCsv(files []*multipart.FileHeader) []map[string]any {
	result := make([]map[string]any, 0)
	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader, wg *sync.WaitGroup) {
			defer wg.Done()
			count := 0
			oFile, err := file.Open()
			if err != nil {
				result = utils.ManageResult(result, file.Filename, "failed", errors.New("file opening error"))
				// errorChan <- err
				return
			}
			defer oFile.Close()
			reader := csv.NewReader(oFile)
			headers, err := reader.Read()

			if err != nil {
				result = utils.ManageResult(result, file.Filename, "failed", errors.New("error while reading the csv file"))
				// errorChan <- err
				return
			}
			validHeader := validHeadersAndPositions(headers)
			fmt.Println(validHeader)

			for {

				data, err := reader.Read()
				count++
				if err != nil {
					break
				}
				if count > 1 {
					for i, field := range data {

						if header, ok := validHeader[i]; ok {
							err := ListofHeadersAndValidations(header, field)
							if err != nil {
								result = utils.ManageResult(result, file.Filename, "validation failed on row "+strconv.Itoa(count), err)
							}
						}
						// err := ListofHeadersAndValidations(headers[i], field)
						// fmt.Println(err)

					}
				}

			}

		}(file, &wg)
		wg.Wait()
		// if err, ok := <-errorChan; ok {
		// 	return err
		// }
		// return nil

	}
	return result
}

func ListofHeadersAndValidations(header, field string) error {

	firstName := func(name string) error {
		if len(name) < 3 {
			return fmt.Errorf("name %s doesnot have characters morethan 3", name)

		}
		return nil
	}
	lastName := func(name string) error {
		if len(name) < 2 {
			return fmt.Errorf("lastname %s must consist atleast 2 characters", name)
		}
		return nil
	}
	headerList := map[string]func(string) error{

		"firstname": firstName,

		"lastname": lastName,
	}

	if val, ok := headerList[strings.ToLower(header)]; ok {

		return val(field)
	}

	return nil

}
func validHeadersAndPositions(headers []string) map[int]string {
	validHeader := make(map[int]string)
	var headersList = map[string]bool{
		"firstname": true,
		"lastname":  true,
	}

	for pos, header := range headers {
		if headersList[strings.ToLower(header)] {
			validHeader[pos] = header
		}
	}
	return validHeader
}
