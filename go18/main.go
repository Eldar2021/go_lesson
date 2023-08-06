package main

import "fmt"

/// Defer (ertelemek)
/*
Defer kelimesinin Türkçe’deki karşılığı ertelemektir.
Bu deyimi yapacağımız işlemin başına eklersek o işlemi
içerisinde bulunduğu fonksiyonun içindeki işlemlerden sonra çalıştırır.
Çok karışık bir cümle kurdum ama uygulamaya geçince anlayacaksınız.
*/

func main() {
	defer fmt.Println("first sentence")
	fmt.Println("second sentence")
	// second sentence, first sentence

	/*
		Açıklamaya gelirsek ekrana İlk Cümle yazısını bastıran satırımızın
		başına defer terimini ekledik. `defer` eklediğimiz satır `main()`
		fonksiyonunun içinde olduğu için `main()` fonsyionundaki tüm işlemler
		tamamlandıktan sonra ekrana yazımızı bastırdı.

		Birden fazla defer ekleyecek olursak;
	*/
	defer fmt.Println("ilk Cümle")
	defer fmt.Println("İkinci Cümle")
	defer fmt.Println("Üçüncü Cümle")
	defer fmt.Println("Dördüncü Cümle")
	fmt.Println("Beşinci Cümle")
	// Beşinci Cümle, Dördüncü Cümle, Üçüncü Cümle, İkinci Cümle, ilk Cümle

	/*
		Burdan anlıyoruz ki en baştaki defer eklenen satır en son işleme
		tabi tutuluyor. Hadi defer ile alakalı bir programlama alıştırması
		yapalım.
	*/

	fmt.Println("Sayıyor")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Bitti")
}
