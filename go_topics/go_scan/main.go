package main

import (
	"bufio"
	"fmt"
	"os"
)

/// Golang'te Kullanıcıdan Giriş Alma
/*
* Golang’te diğer programlama dillerinde de olduğu gibi kullanıcıdan değer
* girişi alınabilir. Böylece programımızı interaktif hale getirmiş oluruz.
 */
func main() {
	// `Scan()` Fonksiyonu
	/*
	* Bu fonksiyon boşluğa kadar olan kelimeyi kaydeder.
	* Yeni satır boşluk olarak sayılır. Kullanımını görelim
	 */
	var yazi string
	fmt.Scan(&yazi) // yazi değişkenine değer girilmesini istedik.
	fmt.Println("\n" + yazi)

	// `Scanf()` Fonksiyonu
	/*
	* Scanf() fonksiyonu Printf() fonksiyonu gibi format içerir.
	* Bu fonksiyon ile kullanıcının girişini bölüp birkaç değişkene
	* kaydedebiliriz. Hemen kullanımını görelim.
	 */
	var kelime1, kelime2 string
	fmt.Scanf("%s %s", &kelime1, &kelime2)
	fmt.Println(kelime1)
	fmt.Println(kelime2)

	// Reader ile Satır Olarak Değer Alma
	/*
	* Aşağıdaki yöntem ile bir satır yazıyı giriş olarak alabilirsiniz.
	 */
	giris := bufio.NewReader(os.Stdin)
	str, _ := giris.ReadString('\n')
	fmt.Println(str)
}
