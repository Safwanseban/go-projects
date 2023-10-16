package controllers

import (
	"csv-handle/controllers/interfaces"
	usecase "csv-handle/usecases/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	result := c.csvUsecases.ValidateCsv(files)

	if len(result) > 0 {
		ctx.JSON(http.StatusMultiStatus, gin.H{
			"message": "some validation issues in csv files",
			"result":  result,
		})
		return
	}
	results, count := c.csvUsecases.BatchUpload(files)
	if count >= 1 {
		ctx.JSON(http.StatusMultiStatus, gin.H{
			"message": strconv.Itoa(count) + " files are not completed uploading",
			"results": results,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "uploaded all files successfully",
	})

}
