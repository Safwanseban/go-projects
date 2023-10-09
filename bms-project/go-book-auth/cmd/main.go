package main

import (
	"github.com/Safwanseban/go-book-autha/internal/di"
)

func main() {

	di.InitializeServer().Start()
}
