package main

import (
	"fmt"
	"os"
)

/// Komut Satırı Argümanları (Args)
/*
* Golang ile programlarımızın komut satırı üzerinden
* argümanlar ile çalışmasını sağlayabiliriz
 */

// run: go run . --eldi

func main() {
	// `os` paketimizdeki Args fonksiyonu bize string dizi sunar.
	for i, arg := range os.Args {
		fmt.Println(i, "->", arg)
	}
}
