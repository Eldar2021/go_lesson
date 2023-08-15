package main

import (
	"fmt"
	"os"
)

/// Bir Dizindeki Dosya ve Klasörleri Sıralama
/*
* Golang üzerinde adresini belirlediğimiz bir dizindeki
* dosya ve klasörleri listelemeyi göreceğiz.
 */

func readFolder(d string) {
	dir, err := os.Open(d)
	if err != nil {
		fmt.Println("Dizrectery not dound!")
		os.Exit(1)
	}

	defer dir.Close()

	/*
	* dizin.Readdirnames(0) diyerek tüm dosya ve klasörleri bu
	* değişkenimizin içerisine attık. Burada sıfır kullanmamızın sebebi
	* tüm dosya ve klasörleri okuyabilmek içindir.
	 */
	list, _ := dir.Readdirnames(0)

	for _, isim := range list {
		fmt.Println(isim)
	}

	fmt.Printf("Toplamda %d içerik bulundu.\n", len(list))
}

func main() {
	readFolder("../")
}
