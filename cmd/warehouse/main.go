package main

import (
	"fmt"
	"os"

	"github.com/lvlBA/restApi/internal/app"
)

func main() {
	cfg := app.Config{
		ListenAddress:  ":8080",
		ListenAddress2: ":80",
	}

	if err := app.Run(&cfg); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
