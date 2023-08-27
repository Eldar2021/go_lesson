package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/// Combining Short Options
/*
* Tabii ki. Go urfave/cli paketindeki Combining Short Options, kısa seçenekleri
* birleştirmenin bir yoludur. Kısa seçenekler, -f veya --flag gibi bir tek karakter
* veya bir karakter dizisi olabilir.
*
* Kısa seçenekleri birleştirmek, komut satırının daha okunaklı ve kompakt olmasını
* sağlayabilir. Örneğin, -f -g yerine -fg kısa seçenekleri birleştirilebilir.
*
* Kısa seçenekleri birleştirmek için, UseShortOptionHandling özelliğini true olarak
* ayarlayabilirsiniz. Bu, NewApp() metodunu çağırırken yapılır.
*
* İşte UseShortOptionHandling özelliğini kullanarak kısa seçenekleri birleştirme örneği:
 */

func main() {
	app := &cli.App{
		Name:                   "my-app",
		Usage:                  "My simple application",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:    "hello",
				Aliases: []string{"h"},
				Usage:   "Say hello",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
					&cli.StringFlag{Name: "lastname", Aliases: []string{"l"}},
					&cli.IntFlag{Name: "age", Aliases: []string{"a"}},
				},
				Action: func(c *cli.Context) error {
					fmt.Printf("Hi! My name is %s, My last name is %s, I'm %d years old.\n", c.String("name"), c.String("lastname"), c.Int("age"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
