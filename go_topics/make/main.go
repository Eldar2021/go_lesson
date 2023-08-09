package main

import "fmt"

/// make()
/*
* Go dilinde, make() işlevi, belirli bir türden (slice, map veya channel gibi)
* dinamik olarak boyutlandırılabilir veri yapılarını oluşturmak için kullanılır.
* Bu işlev, bu veri yapılarının dinamik olarak büyütülmesini ve yönetilmesini sağlar.
*
* make() işlevinin genel sözdizimi şu şekildedir:
* `make(tür, uzunluk, kapasite)`
* 1) tür: Oluşturulacak veri yapısının türünü belirtir (slice, map veya channel).
* 2) uzunluk: Veri yapısının başlangıç uzunluğunu belirtir (slice ve channel için geçerlidir).
* 3) kapasite: Veri yapısının başlangıç kapasitesini belirtir (sadece slice için geçerlidir).
 */

func main() {
	slice := make([]int, 5, 10)
	myMap := make(map[string]int)
	ch := make(chan int)

	fmt.Println(slice)
	fmt.Println(myMap)
	fmt.Println(ch)
}

/*
* make() işlevi, belirli veri yapılarını oluştururken başlangıç uzunluğu veya
* kapasitesini belirtmeye olanak tanır. Bu, bellek kullanımını optimize etmek
* ve veri yapılarının performansını artırmak için önemlidir. Ayrıca, make()
* işlevi kullanılarak oluşturulan veri yapılarının türleri ve özellikleri
* derleme zamanında belirlenir, bu da tür güvenliği sağlar.
*
* Not: make() işlevi sadece slice, map ve channel türlerinde kullanılır.
* Diğer türler için, new() veya ilgili türün değeri kullanılarak örnekler
* oluşturulabilir.
 */
