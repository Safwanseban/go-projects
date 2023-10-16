package routes

import (
	"csv-handle/controllers/interfaces"

	"github.com/gin-gonic/gin"
)

type AppServer struct {
	engine *gin.Engine
}

func NewAppServer(csvhandler interfaces.CsvControllerI) *AppServer {

	engine := gin.Default()
	engine.POST("/upload-files", csvhandler.UploadController)
	return &AppServer{
		engine: engine,
	}
}
func (srv *AppServer) Start() {
	if err := srv.engine.Run(":3000"); err != nil {
		panic("error initializing server")
	}
}
