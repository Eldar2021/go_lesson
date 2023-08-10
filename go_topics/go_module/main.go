package main

import (
	"fmt"
	"gomodule/mypackage"

	"github.com/spf13/cobra"
)

/// Go Modules
/*
* Go Modules, Go programlarında kullanılan bir bağımlılık yönetim
* sistemidir. Go Modules, Go programlarının hangi sürümleri
* kullanacağını belirlemek ve Go programlarını derlemek için kullanılır.
 */

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

			mypackage.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}

/*
* Go Modules, Go 1.11 sürümünde tanıtıldı. Go Modules,
* Go programlarının daha güvenilir ve tekrarlanabilir olmasını sağlar.
*
* Go Modules, aşağıdaki gibi kullanılır:
*   - go mod init komutu, yeni bir Go projesi için bir modül oluşturur.
*   - go mod tidy komutu, modülün bağımlılıklarını günceller.
*   - go mod vendor komutu, modülün bağımlılıklarını yerel bir dizinde kopyalar.
*   - go mod download komutu, modülün bağımlılıklarını indirir.
*   - go mod graph komutu, modülün bağımlılık grafiğini görüntüler.
*   - go mod edit komutu, modülün modül dosyasını düzenler.
*   - go mod verify komutu, modülün bağımlılıklarının doğruluğunu doğrular.
*
* Go Modules, Go programlarının daha güvenilir ve tekrarlanabilir
* olmasını sağlar. Go Modules, Go programlarının bağımlılıklarını
* yönetmeyi kolaylaştırır ve Go programlarının derleme sürecini hızlandırır.
 */
