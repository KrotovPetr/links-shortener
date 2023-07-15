package main

import (
	"fmt"
	"github.com/KrotovPetr/links-shortener.git/internal/app"
)

func main() {
	parseFlags()

	fmt.Println("Running server on", flagRunAddr)

	if err := app.Run(flagRunAddr); err != nil {
		panic(err)
	}

}
