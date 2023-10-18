package interfaces

import "mime/multipart"

type CsvUsecase interface {
	BatchUpload(files []*multipart.FileHeader) ([]map[string]any, int)
	ValidateCsv(files []*multipart.FileHeader) ([]map[string]any, []map[string]any)
}
