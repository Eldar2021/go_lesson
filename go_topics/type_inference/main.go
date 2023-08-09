package main

import "fmt"

/// Type Inference
/*
* "Type Inference" (Tür Çıkarımı), Go dilinin bir özelliğidir ve değişken
* türlerini açıkça belirtmek yerine, derleyicinin değişkenin türünü otomatik
* olarak çıkarabilmesine olanak tanır. Bu, kodun daha temiz ve okunabilir
* olmasına yardımcı olurken aynı zamanda esnekliği de sağlar.
*
* Go dilinde tür çıkarımı, genellikle değişkenler tanımlanırken ve fonksiyonlar
* çağrılırken kullanılır. Derleyici, değişkenin veya fonksiyonun türünü,
* sağlanan değere veya dönüş değerine bakarak otomatik olarak belirler.
 */

// Değişkenlerde Tür Çıkarımı:
func main1() {
	/*
	* Yukarıdaki örnekte, age değişkeni için tür belirtilmemiştir ancak
	* derleyici, sağlanan değere bakarak otomatik olarak int türünü çıkarır.
	* Aynı şekilde, name değişkeni için de string türü otomatik olarak çıkarılır.
	 */
	age := 25       // int türünde bir değişken
	name := "Ahmet" // string türünde bir değişken

	fmt.Printf("Yaş: %d\n", age)
	fmt.Printf("İsim: %s\n", name)
}

// Fonksiyonlarda Tür Çıkarımı:
func add(a, b int) int {
	return a + b
}

func main2() {
	/*
	* Yukarıdaki örnekte, add fonksiyonu için parametre ve dönüş değeri
	* türleri açıkça belirtilmemiştir. Ancak derleyici, fonksiyonun
	* çağrıldığı yerdeki argümanlara ve dönüş değerine bakarak türleri
	* otomatik olarak çıkarır.
	 */
	result := add(5, 3)
	fmt.Println("Toplam:", result)
}

// run all code
func main() {
	main1()
	main2()
}

/*
* "Type Inference" özelliği, Go dilinin kodunun daha sade ve esnek olmasını
* sağlar. Ancak bazen açık tür belirtmenin, kodun okunabilirliğini artırabileceği
* veya tür hatalarını daha erken tespit edebileceği durumlar olabilir.
* Bu nedenle, tür çıkarımını kullanırken dikkatli olmak önemlidir.
 */
