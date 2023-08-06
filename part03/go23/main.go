package main

import "fmt"

/// Diziler (Arrays)
/*
Diziler içlerinde bir veya birden fazla değer tutabilen birimlerdir.
Bir dizideki her değer sırasıyla numaralandırılır.
Numaralandırma sıfırdan başlar. Aynı şekilde örneğe geçelim.
*/

func main() {
	// [] array
	// 3 max length array
	// string type array
	var a [3]string
	a[0] = "Ayşe"     // Birinci değer
	a[1] = "Fatma"    // İkinci değer
	a[2] = "Hayriye"  // Üçüncü değer
	fmt.Println(a)    // Çıktımız: [Ayşe Fatma Hayriye]
	fmt.Println(a[1]) // Çıktımız: Fatma

	// ------------------------------------------------
	// {"Ayşe", "Fatma", "Hayriye"} values array
	b := [3]string{"Ayşe", "Fatma", "Hayriye"}
	fmt.Println(b) //Çıktımız: [Ayşe Fatma Hayriye]
}
