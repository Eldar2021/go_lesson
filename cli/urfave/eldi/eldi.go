package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "eldi",
		Usage: "hi eldiiar almazbek",
		Action: func(ctx *cli.Context) error {
			fmt.Printf("Hello %q", ctx.Args().Get(0))
            return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
