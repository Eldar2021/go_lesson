package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

/// Timestamp Flag
/*
* Go urfave/cli paketindeki Timestamp Flag, bir CLI uygulamasında bir
* tarih ve saat değeri almak için kullanılır. Zaman damgası bayrağı,
* uygulamanın saat dilimi ve saat formatına göre tarih ve saat
* değerini döndürür.
*
* Go urfave/cli paketinde, TimestampFlag adı verilen yerleşik bir
* bayrak vardır. Bu bayrak, -t veya --timestamp kısaltmalarıyla
* kullanılabilir.
 */

func main() {
	app := &cli.App{
		Name:  "e-time",
		Usage: "My simple application",
		Flags: []cli.Flag{
			&cli.TimestampFlag{
				Name:   "meetup",
				Usage:  "Meetup Timestamp",
				Layout: "2006-01-02T15:04:05",
				Action: func(ctx *cli.Context, t *time.Time) error {
					fmt.Printf("meetup %s", ctx.Timestamp("meetup").String())
					return nil
				},
			},
			&cli.TimestampFlag{
				Name:     "start",
				Usage:    "start time",
				Layout:   "2006-01-02T15:04:05",
				Timezone: time.Local,
				Action: func(ctx *cli.Context, t *time.Time) error {
					fmt.Printf("start %s", ctx.Timestamp("start").String())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
