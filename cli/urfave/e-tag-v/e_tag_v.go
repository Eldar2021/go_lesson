package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/// Version Flag
/*
* Go urfave/cli paketindeki Version Flag, bir CLI uygulamasının
* sürümünü almak için kullanılır. Sürüm bayrağı, uygulamanın sürümünü
* ve diğer bazı bilgileri görüntüler.
*
* Go urfave/cli paketinde, VersionFlag adı verilen yerleşik bir bayrak
* vardır. Bu bayrak, -v veya --version kısaltmalarıyla kullanılabilir.
 */

func main() {
	app := cli.NewApp()
	app.Name = "Version Flag"
	app.Usage = "Version Flag"
	app.Version = "0.0.3"

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
