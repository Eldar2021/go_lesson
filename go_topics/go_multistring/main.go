package main

import "fmt"

/// Çok Satırlı String Oluşturma
/*
* Çok satırlı string oluşturmak için (`) işaretini kullanırız. Türkçe klavyeden
* alt gr ve virgül tuşuna basarak bu işareti koyabilirsiniz. İşte örnek kodumuz;
 */
func main() {
	yazi := `Bu bir
    çok satırlı
    yazı örneğidir.
`
	fmt.Printf(yazi)
	fmt.Printf("-----------------------------------")
	fmt.Printf("%s", yazi)
}
