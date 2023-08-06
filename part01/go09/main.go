package main

import "fmt"

/// Kod Gruplama İşlemi
/*
Kod gruplama işlemi çok basit bir işlemdir.
Bu işlem sayesinde aynı objeler bloklara göre farklı çalışabilir.
Kodları gruplama için süslü parantez kullanırız. Örneğimizi görelim.
*/

func main() {
	value01 := "bir"
	{
		value01 := "iki"
		fmt.Println(value01)
	}
	fmt.Println(value01)

	// -----------------------------------------------------------------
	// & (and) işareti ile değişkenin bellekteki adresini öğrenebiliriz.
	// -----------------------------------------------------------------

	value03 := "uch"
	{
		value04 := "dort"
		fmt.Println(value04)
		fmt.Println(&value04) // 0x1400008e250
	}
	fmt.Println(value03)
	fmt.Println(&value03) // 0x1400008e240
}
