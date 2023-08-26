package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

/// Subcommands Categories
/*
* `urfave/cli`'deki Subcommands Categories, alt komutları mantıksal
* kategorilere ayırmanın bir yoludur. Bu, kullanıcıların alt komutları
* daha kolay bulmalarına yardımcı olabilir.
*
* Subcommands Categories, Categories() metodunu kullanarak bir uygulamaya
* eklenir. Categories() metodu, bir []string türündeki bir dizi kategori adı alır.
 */

func main() {
	app := cli.NewApp()
	app.Name = "my-app"
	app.Usage = "A simple CLI application with subcommands"

	app.Commands = []cli.Command{
		{
			Name:     "edit",
			Category: "file",
		},
		{
			Name:     "read",
			Category: "file",
		},
		{

			Name:     "create",
			Category: "file",
		},
		{
			Name:     "create",
			Category: "folder",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
