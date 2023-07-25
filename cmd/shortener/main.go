package main

import (
	"fmt"
	"github.com/KrotovPetr/links-shortener.git/cmd/shortener/config"
	"github.com/KrotovPetr/links-shortener.git/internal/app"
	"log"
)

func main() {
	configVar := config.GetConfig()

	fmt.Println("Running server on", configVar.Address)

	if err := app.Run(configVar.Address); err != nil {
		log.Fatal(err)
	}

}
