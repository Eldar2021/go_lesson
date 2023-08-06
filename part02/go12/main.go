package main

import "fmt"

/// Fonksiyon Çeşitleri
// Golang’ta genel olarak 3 çeşit fonksiyon yapısı bulunmaktadır.
// Hemen bu çeşitleri görelim.

/// Variadic Fonksiyonlar
// Variadic fonksiyon tipi ile fonksiyonumuza kaç tane değer girişi
// olduğunu belirtmeden istediğiniz kadar değer girebilirsiniz.

/*
Yukarıdaki fonksiyonumuzu inceleyelim. Vereceğimiz sayıları toplaması
için aşağıda toplama adında bir fonksiyon oluşturduk. Fonksiyonun
parametresi içerisine, yani parantezler içerisine, sayilar isminde
int tipinde bir değişken tanımladık. … (üç nokta) ile istediğimiz
kadar değer alabileceğini belirttik. toplam değerini mantıken doğru
değer vermesi için 0 yaptık. Çünkü her sayıyı toplam değikeninin
üzerine ekleyecek.

range’in buradaki kullanım amacından bahsedeyim. range’i for döngüsü
ile kullandığımızda işlem yaptığımız öğenin uzunluğuna göre işlemimizi
sürdürürüz. Yani fonksiyonumuzun içine ne kadar sayı eklersek işlemimiz
ona göre şekillenecektir. For ve Range işlemini daha sonraki bölümümüzde
göreceğiz.

Range kullanımında _, n şeklinde değişken tanımlamamızın sebebi, birinci
değişken yani _, dizinin indeksini yani sıra numarasını verir. Bizim
bununla bir işimiz olmadığı için _ koyarak kullanmayacağımızı belirttik.
İkinci değişken ise yani n dizinin içindeki değeri verir yani fonksiyona
girdiğimiz sayıları. Sonuç olarak bu fonksiyonda return ile for işleminden
sonra tüm sayıların toplamını döndürüp main() fonksiyonu içerisinde
ekrana bastırmış olduk.
*/

func kosh(sandar ...int) int {
	summa := 0
	for _, v := range sandar {
		summa += v
	}
	return summa
}

/// Recursive (İç-içe) Fonksiyonlar
// Recursive fonksiyonlar yazdığımız fonksiyonun içinde aynı
// fonksiyonu kullanmamız demektir. Fonksiyonumun tüm işlemler
// bittiğinde return olur. Örneğimize geçelim.

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

/// Closure (Atanmış) Fonksiyonlar
// Closure fonksiyonlar ile değişkenlerimizi fonksiyon olarak
// tanımlayabiliriz. Buradan anlamamış gereken şey fonksiyonların
// da atanabilen veri tipleri olduğudur.

func customPrint() func(s string) {
	return func(s string) {
		fmt.Println(s)
	}
}

func main() {
	fmt.Println(kosh(1, 2, 5, 6, 6))

	summa := func(sandar ...int) int {
		summa := 0
		for _, v := range sandar {
			summa += v
		}
		return summa
	}

	fmt.Println(summa(12, 3, 7))

	result := customPrint()
	result("Eldi")

	fmt.Println(factorial(6))

}
