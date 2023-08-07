package main

import "fmt"

/// Dilimler (Slices)
/*
Dilimler bir dizideki değerlerin istediğimiz bölümünü kullanmamıza yarar.
Yani diziyi bir pasta olarak düşünürsek kestiğimiz dilimi yiyoruz sadece.
*/

func main() {
	// -------------------------------------------------------------------------
	// Simple Dilim
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a) // [1 2 3 4 5 6]
	b := a[2:4]
	fmt.Println(b) // [3 4]
	/*
	* İnceleme kısmına geçelim. a isminde 6 tane int tipinde değer
	* alan bir dizi oluşturduk. Çıktımızın içeriğini görmek için ekrana
	* bastırdık. Dilimleme işlemi olarak yorum yaptığım satırda ise a
	* dizisinde 2 ve 4 indeksi arasındaki değerleri dizi olarak b’ye
	* kaydettik. b dizisinin içeriğini ekrana bastırdığımızda ise dilimlenmiş
	* alanımızı gördük. Dilimleme işleminde [ ] içerisine dilimlemenin
	* başlayacağı ve biteceği indeksi yazarız.
	 */

	// -------------------------------------------------------------------------
	// Dilim Varsayılanları (Sıfır Değerleri)
	c := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(c) // [1 2 3 4 5 6 7]
	var d []int = c[:4]
	fmt.Println(d) // [1 2 3 4]
	e := c[4:]
	fmt.Println(e) // [5 6 7]
	// f := c[2:9] give error (not index)

	// -------------------------------------------------------------------------
	// Dilim Uzunluğu ve Kapasitesi
	g := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(g) // [1 2 3 4 5 6]
	h := g[3:6]
	fmt.Println(h)                    // [4 5 6]
	fmt.Println("g uzunluk", len(g))  // g uzunluk 6
	fmt.Println("h uzunluk", len(h))  // h uzunluk 3
	fmt.Println("g kapasite", cap(g)) // g kapasite 6
	fmt.Println("h kapasite", cap(h)) // h kapasite 3
	/*
	* `h` dizisi ile `g` dizisini dilimlediğimiz için `h` dizisinin kapasitesi ve
	* uzunluğu değişti. Uzunluk `len()` dizinin içindeki değerlerin sayısıdır.
	* Kapasite `cap()` ise dizinin maksimum alabileceği değer sayısıdır.
	 */

	// -------------------------------------------------------------------------
	//  Boş Dilimler (Nil Slices)
	/*
	* Boş bir dilimin varsayılan (sıfır) değeri nil’dir. Örnek olarak;
	 */

	var i []int
	fmt.Println(i) // []
	if i == nil {  // nil == dart null || empty
		fmt.Println("Boş")
	}

	// -------------------------------------------------------------------------
	// Make ile Dilim Oluşturma
	// `0` uzunluğu vermeye de biliriz
	// `5` kapasitesi
	j := make([]int, 0, 5)
	fmt.Println(j) // []

	// -------------------------------------------------------------------------
	// Dilime Ekleme Yapma

	var k []int
	fmt.Println(k) // []
	k = append(k, 1)
	fmt.Println(k) // [1]
	k = append(k, 5)
	fmt.Println(k)                    // [1 5]
	fmt.Println("k uzunluk", len(k))  // h uzunluk 2
	fmt.Println("k kapasite", cap(k)) // g kapasite 2
}
