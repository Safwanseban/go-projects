package main

import (
	"github.com/Safwanseban/go-book-gateway/internal/di"
)

func init() {

}
func main() {

	appServer := di.InitializeApi()
	appServer.Start()

}
