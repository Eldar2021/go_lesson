package main

import "fmt"

// Deferred Functions Calls
/*
* Go, bir fonksiyon çağrısını fonksiyon tamamlandıktan sonra çalıştırılmak
* üzere zamanlayan defer adlı özel bir deyime sahiptir.
 */

func main() {
	defer fmt.Println("Second")
	fmt.Println("First")

	/*
	* Bunun çeşitli avantajları vardır:
	* 1) Close çağrımızı Open çağrımıza yakın tutar, böylece anlaşılması daha kolay olur.
	* 2) Fonksiyonumuzun birden fazla geri dönüş ifadesi varsa (belki bir if ve bir else içinde), Close her ikisinden de önce gerçekleşecektir.
	* 3) Ertelenmiş Fonksiyonlar, çalışma zamanı paniği oluşsa bile çalıştırılır.
	* 4) Ertelenmiş fonksiyonlar LIFO sırasına göre çalıştırılır, bu nedenle yukarıdaki kod yazdırılır: 4 3 2 1 0.
	* 5) Bu örnekte olduğu gibi "ertelenmiş listeye" birden fazla fonksiyon koyabilirsiniz.
	 */

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
