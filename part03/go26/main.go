package main

import "fmt"

/// Map
/*
`Map`’in Türkçe karşılığında yapacağı işlemi anlatan bir çeviri
olmadığı için anlamı yerine yaptığı işi bilelim. Map ile bir değişken
içerisindeki dizileri bölge olarak ayırabiliriz
*/
/*
Go dilinde "map" kavramı, verileri anahtar-değer çiftleriyle ilişkilendiren
bir veri yapılarından biridir. Diğer dillerde "dictionary", "hash table"
veya "associative array" olarak da bilinir. Go'da map, aynı türden anahtarları
aynı türden değerlerle eşleştirmek için kullanılır.
*/

// Go dilinde map, aşağıdaki şekilde tanımlanır:
// `map[keyTipi]degerTipi`

type persons struct {
	p1, p2, p3 string
}

func main() {
	m := make(map[string]persons)
	m["name"] = persons{"Eldi", "Ruslan", "Kylychnek"}
	fmt.Println(m["name"]) // {Eldi Ruslan Kylychnek}

	// Bir string tipinde anahtarları ve integer tipinde değerleri olan bir map tanımlayalım
	a := make(map[string]int)

	// Anahtar-değer çiftlerini map'e ekleyelim
	a["Eldi"] = 12
	a["Ruslan"] = 23
	a["Kylych"] = 34

	fmt.Println(a["Eldi"]) // 12

	// Map'teki bir değeri silmek için "delete" fonksiyonunu kullanabiliriz
	delete(a, "Eldi")
	fmt.Println(a["Eldi"]) // 0

	// Belirli bir anahtarın map içinde olup olmadığını kontrol edebiliriz
	x, y := a["Eldi"]
	fmt.Println(x, y) // 0, false
	if !y {
		fmt.Println("Eldi is not find")
	}
}
