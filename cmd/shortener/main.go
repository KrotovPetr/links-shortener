package main

import (
	"fmt"
	"github.com/KrotovPetr/links-shortener.git/cmd/shortener/config"
	"github.com/KrotovPetr/links-shortener.git/internal/app"
)

var addressFlag, urlFlag string

func main() {
	config := config.GetConfig()

	fmt.Println("Running server on", config.Address)

	if err := app.Run(config.Address); err != nil {
		panic(err)
	}

}
