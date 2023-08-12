package main

import "fmt"

/// Sprintf
/*
* Sprintf fonksiyonu fmt paketine dahil bir fonksiyondur.
* Bu fonksiyon değişkenlere formatlı atama yapmamıza yardımcı olur.
 */

func main() {
	isim := "Kaan"

	isimTip := fmt.Sprintf("isim değişkeni %T tipindedir.", isim)

	fmt.Println(isimTip)
}
