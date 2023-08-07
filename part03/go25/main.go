package main

import "fmt"

/// Range
/*
Range, üzerinde kullanıldığı diziyi for döngüsü ile tekrarlayabilir.
Bir dilim range edildiğinde, tekrarlama başına iki değer döndürür (return).
Birinci değer dizinin indeksi, ikinci değer ise bu indeksin içindeki değerdir.
*/

func main() {
	persons := [4]string{"Eldi", "Rus", "Kylych"}
	fmt.Println(persons) // [Eldi Rus Kylych ]

	for a, b := range persons {
		fmt.Println(a, b) // 0 Eldi, 1 Rus, 2 Kylych, 3
	}

	/*
	* Yukarıdaki yazdığımız kodları açıklayalım. isimler isminde içerisinde
	* string tipinde değerler olan bir dizi oluşturduk.

	* For döngümüz ile dizinimizdeki değerleri sıralayacak bir sistem oluşturduk.
	* Döngümüzü açıklayacak olursak, bahsettiğimiz gibi dizi üzerinde uygulanan
	* range terimi iki değer döndürecek olduğundan bu değerleri kullanabilmek
	* için a ve b adında argüman belirledik. range isimler diyerek isimler
	* dizisini kullanacağımızı belirttik. Ekrana bastırma bölümümüzde ise %
	* işaretleri ile sağ taraftan hangi değerleri nerede kullanacağımızı belirttik.
	 */
}
