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

func (cu *CsvUsecases) ValidateCsv(files []*multipart.FileHeader) ([]map[string]any, []map[string]any) {
	result := make([]map[string]any, 0)
	fileStatus := make([]map[string]any, 0)
	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader, wg *sync.WaitGroup) {
			defer wg.Done()
			oFile, err := file.Open()
			if err != nil {
				result = utils.ManageResult(result, file.Filename, "failed", errors.New("file opening error"))

				return
			}
			defer oFile.Close()
			reader := csv.NewReader(oFile)
			headers, err := reader.Read()

			if err != nil {
				result = utils.ManageResult(result, file.Filename, "failed", errors.New("error while reading the csv file"))

				return
			}
			errorCount := 0
			validHeader := validHeadersAndPositions(headers)
			fmt.Println(validHeader)
			row := 1
			for {
				data, err := reader.Read()
				row++
				if err != nil {
					break
				}

				for i, field := range data {

					if header, ok := validHeader[i]; ok {
						err := ListofHeadersAndValidations(header, field)
						if err != nil {
							errorCount++
							result = utils.ManageResult(result, file.Filename, "validation failed on row "+strconv.Itoa(row), err)
						}
					}
				}
			}
			if errorCount == 0 {
				if err := uploadFiles(file, cu); err != nil {
					result = utils.ManageResult(result, file.Filename, "uploading failed", err)
				}
				fileStatus = append(fileStatus, map[string]any{
					"fileName": fmt.Sprintf("%s has been uploaded", file.Filename),
				})
				return
			}
			fileStatus = append(fileStatus, map[string]any{
				"fileName": fmt.Sprintf("%s have some errors ", file.Filename),
			})

		}(file, &wg)
		wg.Wait()
	}
	return result, fileStatus
}

func uploadFiles(file *multipart.FileHeader, cu *CsvUsecases) error {
	fileName := strings.Split(file.Filename, ".")
	extension := fileName[len(fileName)-1]
	ofile, err := file.Open()
	if err != nil {
		return errors.New("file opening error")
	}
	defer ofile.Close()
	if err := cu.Svc.UploadToS3("csvhandlebucket", utils.NewUUIDFile(extension), ofile); err != nil {
		return errors.New("error uploading file")
	}
	return nil
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
	headerList := map[string]any{

		"firstname": firstName,

		"lastname": lastName,
	}

	if val, ok := headerList[strings.ReplaceAll(strings.ToLower(header), " ", "")]; ok {

		return val.(func(string) error)(field)
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
		if headersList[strings.ReplaceAll(strings.ToLower(header), " ", "")] {
			validHeader[pos] = header
		}
	}
	return validHeader
}
