package main

import "fmt"

/// Dinamik Değişkenler
/*
Go dilinde dinamik değişkenler, tipleri derleme zamanında belirlenmeyen
değişkenlerdir. Bu, değişkenin değerine bağlı olarak tipinin değişebileceği
anlamına gelir. Dinamik değişkenler, tipleri önceden bilinmeyen verileri
depolamak için kullanışlıdır.

Dinamik değişkenler, var deyimi kullanılarak oluşturulur. var deyiminin
tip kısmı boş bırakılırsa, değişken dinamik tipli olur.

Aşağıdaki örnekte, x değişkeni dinamik tiplidir:
*/

func main() {
	var x any
	fmt.Println(x) // <nil>
	x = 10
	fmt.Println(x) // 10
	x = "Eldi"
	fmt.Println(x) // Eldi

}
