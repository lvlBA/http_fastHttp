package main

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v7"

	"github.com/lvlBA/restApi/internal/app"
)

func main() {
	cfg := new(app.Config)
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("failed to parse config: %s", err)
		os.Exit(1)
	}

	if err := app.Run(cfg); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
