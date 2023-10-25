package main

import (
	"github.com/Safwanseban/go-book-gateway/internal/di"
)


func main() {

	di.InitializeApi().Start()

}
