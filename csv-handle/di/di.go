package di

import (
	"csv-handle/client"
	"csv-handle/controllers"
	"csv-handle/routes"
	"csv-handle/usecases"
)

func InitializeServer() *routes.AppServer {

	client := client.NewAwsClient()
	csvUsecase := usecases.NewCsvUsecase(client)
	csvController := controllers.NewCsvController(csvUsecase)
	return routes.NewAppServer(csvController)
}
