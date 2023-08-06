package main

import "fmt"

/// Pointers (İşaretçiler)
/*
İşaretçiler ile nesnenin bellekteki adresi üzerinden işlemler yapabilir.
Daha önceden işaretçileri içeren bir dil kullanmamış veya herhangi bir
programlama dili de kullanmamış olabilirsiniz. Bu yüzden daha iyi anlamanız
için işaretçilerin çalışma mantığını bilmemiz gerekir.
*/
/*
Örnek olarak `a` isminde bir değişkenimiz olsun.
Bu değişkenimiz tam sayı `int` tipinde `8` değerini saklıyor olsun.
Bu da Go dilinde aşağıdaki gibi oluyor. `var a int = 8`

Bu değişkenimizi oluşturduktan sonra programımız çalışmaya başlayınca
işletim sistemimiz bu değişkene özel, bellek (RAM) üzerinde bir alan
ayıracaktır. Programın geri kalanında değişkenimizin değerine bu alan
üzerinden ulaşılacaktır. Yani bir değişken oluşturduğumuzda bellek
üzerinde aşağıdaki resimdeki gibi bir alan oluştuğunu hayal edebilirsiniz.
https://go.kaanksc.com/boeluem-2/pointers-isaretciler
Yukarıdaki resimde mor renkli olarak gördüğümüz ifade ise değişkenimizin
bellekteki adresidir. (Bu adres temsilidir. Zaten sürekli olarak
değişen birşeydir.)
*/
// Peki Go'da işaretçileri nasıl kullanırız?
/*
Bir program senaryosu belirleyelim. Bu programımızda yukarıdaki
gibi `a` isminde `int` tipinde `8` değerini tutan bir değişkenimiz olsun.
Ve `a` değişkenimize `5` ekleyen bir fonksiyonumuz olsun.
Son olarak `a`'yı ekrana bastıralım.
*/

func main() {
	a := 8
	add5(a)
	fmt.Println(a) // 8

	/*
		Yukarıdaki örnekte `a` değişkenini ekrana bastırdığımızda sonucun
		hala `8` olduğunu görüyoruz. Halbuki ekle fonksiyonunun içerisinde
		gördüğümün gibi `5` ekliyoruz.

	*/
	/*
		`a` değişkeninin değişmeme sebebi şudur:
		ekle fonksiyonunun parametresi olarak `int` tipinde `v` değişkenini
		oluşturduk. `v` değişkenimiz aslında `a`'dan gelen değeri
		kullanmamızı sağlıyor. Yani bize a'nın kendisini vermiyor.
		`O` yüzden `v`` üzerinde değişiklik yaptığımızda `a`'ya yansımayacaktır.
	*/
	/*
		`a` değişkenini değiştirebilmemiz için bize `a`'nın bellekteki
		adresi gerekiyor. Bunun için de `&` (ampersand) işaretini kullanabiliriz.
	*/

	fmt.Println(&a) // 0x1400018e008
	/*
		Artık `a`'nın bellekteki adresini öğrenebiliyoruz.
		Sıra geldi bu adres üzerinden `a`'nın değerine ulaşabilmeye.
		Bunun için de `*`(yıldız) işaretini kullanabiliriz.
	*/
	b := &a
	fmt.Println(*b) // 8

	// Dogru kullanımı
	x := 5
	fmt.Println(x) // 5
	add5Correct(&x)
	fmt.Println(x) // 10

	y := 8
	fmt.Println(y) // 8
	z := &y
	*z = 10
	fmt.Println(y) // 10

}

func add5(v int) {
	v += 5
}

func add5Correct(v *int) {
	*v += 5
}
