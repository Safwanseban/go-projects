package controllers

import (
	"csv-handle/controllers/interfaces"
	usecase "csv-handle/usecases/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Message      string      `json:"message"`
	FileStatus   interface{} `json:"fileStatus"`
	ErrorResults interface{} `json:"errorResults,omitempty"`
}
type CsvController struct {
	csvUsecases usecase.CsvUsecase
}

func NewCsvController(csvUsecase usecase.CsvUsecase) interfaces.CsvControllerI {

	return &CsvController{
		csvUsecases: csvUsecase,
	}

}

func (c *CsvController) UploadController(ctx *gin.Context) {

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	files := form.File["files"]
	if len(files) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "no files are provided",
		})
		return
	}
	result, uploadedFiles := c.csvUsecases.ValidateCsv(files)
	if len(result) > 0 || len(uploadedFiles) > 0 {
		ctx.JSON(http.StatusMultiStatus, gin.H{
			"message":      "some validation issues in csv files",
			"fileStatus":   uploadedFiles,
			"errorResults": result,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "uploaded all files successfully",
	})

}
