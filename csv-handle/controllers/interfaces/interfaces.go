package interfaces

import "github.com/gin-gonic/gin"

type CsvControllerI interface {
	UploadController(ctx *gin.Context)
}
