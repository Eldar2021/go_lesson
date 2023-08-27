package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/// Exit Codes
/*
* Go urfave/cli paketindeki Exit Codes, bir CLI uygulamasının çalışma durumunu
* belirtmek için kullanılan bir mekanizmadır. Exit kodları, bir uygulamanın bir
* hata nedeniyle başarısız olup olmadığını veya beklendiği gibi çalışıp çalışmadığını
* belirtmek için kullanılabilir.
*
* Exit kodları, bir CLI uygulamasının hata ayıklamasını kolaylaştırmanın bir yoludur.
* Bir uygulama bir hata nedeniyle başarısız olursa, hata ayıklama aracısı, uygulamanın
* çıkış kodunu kullanarak hatanın kaynağını belirlemeye yardımcı olabilir.
*
* Exit kodları, bir CLI uygulamasının durumunu diğer uygulamalara bildirmenin
* bir yoludur. Bir uygulama, bir işlemin tamamlandığını veya bir hata oluştuğunu
* bildirmek için bir çıkış kodu kullanabilir.
*
* Go urfave/cli paketinde, aşağıdakiler dahil olmak üzere çeşitli standart çıkış
* kodları tanımlanmıştır:
*   - 0: Uygulama başarıyla çalıştı.
*   - 1: Uygulama bir hata nedeniyle başarısız oldu.
*   - 2: Bir komut hatası oluştu.
*   - 3: Bir hata ayıklama hatası oluştu.
*   - 127: Uygulama bulunamadı.
*
* Exit kodları, bir CLI uygulamasının Exit() metodunu kullanarak ayarlanır.
* Exit() metodu, bir tam sayı alır ve bu tam sayı, uygulamanın çıkış kodu olur.
 */

func main() {
	app := &cli.App{
		Name:  "my-app",
		Usage: "A simple application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your name",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.String("name") == "" {
				return cli.Exit("name is reqiared", 1)
			}
			fmt.Printf("Hello %s\n", ctx.String("name"))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

/*
macbook_pro@MBPEldiAlmazbek e-exit % e-exit Eldi
name is reqiared
macbook_pro@MBPEldiAlmazbek e-exit % e-exit -name Eldi
Hello Eldi
macbook_pro@MBPEldiAlmazbek e-exit % e-exit -name
Incorrect Usage: flag needs an argument: -name

NAME:
   my-app - A simple application

USAGE:
   my-app [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value  Your name
   --help, -h    show help
flag needs an argument: -name
macbook_pro@MBPEldiAlmazbek e-exit % e-exit -name
Incorrect Usage: flag needs an argument: -name

NAME:
   my-app - A simple application

USAGE:
   my-app [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --name value  Your name
   --help, -h    show help
flag needs an argument: -name
macbook_pro@MBPEldiAlmazbek e-exit % e-exit -name E
Hello E
macbook_pro@MBPEldiAlmazbek e-exit %
*/
