package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/// Bash Completions
/*
* Go urfave/cli paketindeki Bash Completions, bir CLI uygulaması için bash
* tamamlanmasını etkinleştirmek için kullanılır. Bash tamamlanması,
* kullanıcıların komut satırında bir komutun veya seçeneğin tamamını
* yazmasına gerek kalmadan bir komutun seçeneklerini veya parametrelerini
* tamamlamasına olanak tanır.
*
* App nesnesi üzerindeki EnableBashCompletion bayrağını true olarak ayarlayarak
* tamamlama komutlarını etkinleştirebilirsiniz. Varsayılan olarak, bu ayar bir
* uygulamanın alt komutları için otomatik tamamlamaya izin verir, ancak
* Uygulama veya alt komutları için kendi tamamlama yöntemlerinizi de yazabilirsiniz.
 */

func main() {
	app := cli.App{
		Name:                 "e-bash-c",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a",
				Action: func(c *cli.Context) error {
					fmt.Println("added file")
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "remove a",
				Action: func(c *cli.Context) error {
					fmt.Println("removed file")
					return nil
				},
			},
			{
				Name:  "read",
				Usage: "read a file",
				Action: func(c *cli.Context) error {
					fmt.Println("readed file")
					return nil
				},
			},
			{
				Name:  "say",
				Usage: "say name, last name, or age",
				Subcommands: []*cli.Command{
					{
						Name:  "name",
						Usage: "say name",
						Action: func(c *cli.Context) error {
							fmt.Println("Eldiiar")
							return nil
						},
					},
					{
						Name:  "last-name",
						Usage: "say last name",
						Action: func(c *cli.Context) error {
							fmt.Println("Almazbek")
							return nil
						},
					},
					{
						Name:  "age",
						Usage: "say age",
						Action: func(c *cli.Context) error {
							fmt.Println("23")
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
