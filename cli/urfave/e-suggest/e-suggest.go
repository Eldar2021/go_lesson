package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/// Suggestion
/*
* Go urfave/cli paketindeki Suggestions, bir CLI uygulaması için öneriler
* sağlamak için kullanılır. Öneriler, kullanıcıların komut satırında
* doğru seçenekleri ve parametreleri seçmesine yardımcı olabilir.
*
* Go urfave/cli paketinde, Suggestions() metodu kullanılarak öneriler
* sağlanır. Suggestions() metodu, bir komutun seçeneklerini ve
* parametrelerini kabul eder.
 */

func main() {
	app := &cli.App{
		Name:    "e-suggest",
		Usage:   "A CLI application",
		Suggest: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your name",
				Action: func(ctx *cli.Context, name string) error {
					fmt.Printf("My name is %s\n", name)
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "lastname",
				Usage: "Your lastname",
				Action: func(ctx *cli.Context, lastname string) error {
					fmt.Printf("My last name is %s.\n", lastname)
					return nil
				},
			},
			&cli.IntFlag{
				Name:  "age",
				Usage: "Your age",
				Action: func(ctx *cli.Context, age int) error {
					fmt.Printf("I'm %d years old.\n", age)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
