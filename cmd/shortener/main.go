package main

import (
	"fmt"
	"github.com/KrotovPetr/links-shortener.git/cmd/shortener/config"
	"github.com/KrotovPetr/links-shortener.git/internal/app"
)

func main() {
	configVar := config.GetConfig()

	fmt.Println("Running server on", configVar.Address)

	if err := app.Run(configVar.Address); err != nil {
		panic(err)
	}

}
