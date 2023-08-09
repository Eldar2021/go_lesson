package main

import "fmt"

/// Panic and Recover
/*
* Panik ve İyileşme
* Go'da (Golang) panik, beklenmedik veya kurtarılamaz bir durum meydana
* geldiğinde bir programın normal yürütülmesini durdurmanıza olanak tanıyan
* bir mekanizmadır. Diğer programlama dillerindeki istisnalara benzer.
* Bir panik tetiklendiğinde, program çağrı yığınını çözmeye başlar, her işlev
* çağrısıyla ilişkili ertelenmiş işlevleri çalıştırır ve ardından paniğin
* nedenini açıklayan bir günlük mesajıyla çöker.
 */

func main() {
	//---------------------1-------------------------
	f()
	fmt.Println("Returned normally from f.")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panik durumu:", r)
		}
	}()

	//---------------------2-------------------------
	/*
	* Bu örnekte, doPanic fonksiyonu bir panik durumu oluşturuyor. main
	* fonksiyonunda bir defer ifadesi kullanarak bir anonim fonksiyon tanımlıyoruz.
	* Bu anonim fonksiyon recover ile panik durumunu yakalar ve mesajını ekrana
	* yazdırır. Ancak bu yakalama işlemi sonrasında program normal bir şekilde devam
	* eder.
	 */
	fmt.Println("Program başladı")
	doPanic()
	fmt.Println("Bu satır çalışmayacak")
}

func f() {
	defer func() {
		/*
		* Eğer bir panik durumu varsa, recover fonksiyonu panik durumunu yakalar
		* ve durumunun değerini kontrol eder. Eğer r değeri nil değilse, yani
		* bir panik durumu oluştuysa, belirtilen anonim fonksiyon çalıştırılır
		* ve "Recovered in f" mesajını ve panik durumunu yazdırır.
		 */
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func doPanic() {
	panic("bir panik durumu oluştu")
}
