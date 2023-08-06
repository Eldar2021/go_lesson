package main

import "fmt"

/// Switch
/*
`Switch` kelimesinin Türkçe’deki anlamı anahtardır.
`Switch` deyimi de `if-else` deyimi gibi koşul üzerine çalışır.
Yine teorik kısmı geçip anlaşılır olması için örnek yapalım.
`case` deyimi durumu ifade eder. Koşul sağlandığı zaman işleme devam edilmez.
*/

func main() {
	// ------------------------------------------------------------------
	a := 3

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	} //  3

	// ------------------------------------------------------------------
	// Switch’te koşulların gerçekleşmediği zaman işlem uygulamak
	// istiyorsak bunu default terimi ile yaparız. Örnek;
	b := 7

	switch b {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default: // (ech bir sharty tuura kelbese)
		fmt.Println("Bilbeim")
	} //  Bilbeim

	// ------------------------------------------------------------------
	// Koşulsuz Switch
	/*
		Switch’in tanımını daha iyi anlayabilmeniz için koşulsuz
		switch kullanımına örnek verelim. Bu yöntemde switch deyiminin
		yanına koşul girmek yerine case deyiminin yanına koşul giriyoruz.
	*/

	c := 100

	switch {
	case c < 0:
		fmt.Println("Error c < 0")
	case c > 120:
		fmt.Println("Error c > 120")
	case c == 100:
		fmt.Println("Wow")
	case c < 1:
		fmt.Println("Hi, Wellcome to our wolrd")
	} //  Wow

	// ------------------------------------------------------------------
	// Sonraki Koşulu Kontrol Ettirme
	/*
		Durumlar içerisinde kontrol etmemiz gereken başka durumlarda olabilir.
		Bunun için `fallthrough` deyimini kullanabiliriz.
	*/

	d := 5
	switch {
	case d == 5:
		fmt.Println("Yees!")
		fallthrough // dagy ulant(bashkalaryn da karap chyk)
	case d < 10:
		fmt.Println("d 10'dan küçüktür")
	} // Yees!, d 10'dan küçüktür

	// ------------------------------------------------------------------
	// Switch'e Özel Değişken Tanımlama
	/*
		Tıpkı If deyimindeki Switch içerisinde kullanabileceğimiz
		değişkenler tanımlayabiliriz.
	*/

	switch e := 5; {
	case e == 5:
		fmt.Println("e 5'tir")
	case e < 10:
		fmt.Println("e 10'dan küçüktür")
	} // e 5'tir

}
