package main

import (
	"fmt"
	"regexp"
)

/// Regexp (Kurallı İfadeler)
/*
* Regular Expressions (Regexp), modern programlama dillerinin neredeyse hepsinde
* bulunan metinsel ifadelerinizin yapısını kontrol etmenizi sağlayan bir pakettir.
*
* Bu paket sayesinde yazdığımız programda ifadelerin uygunluğunu kontrol edebilir,
* işimize yarayacak ifade/ifadeleri daha kolay ayırabilir ve giriş yapılan
* ifadeleri uygun bir düzene koyabiliriz.
 */

func main() {
	/*
	* nameControl adında bir değişken oluşturduk ve bu değişkenimizde
	* MustCompile() fonksiyonu ile metinsel ifademizin kurallarını belirledik.
	*
	* Belirlediğimiz kural ise dizgimizin küçük harflerle a'dan z'ye kadar ve
	* ek olarak 0'dan 9'a kadar "rakamsal" ifade alabileceğidir. Bu kurala uyulması
	* için metinsel ifade ne eksik ne fazla hiçbir şey olmaması gerekir.
	 */
	nameControl := regexp.MustCompile(`^[a-z]+[0-9]$`)

	// "kaan10" ifadesinde sadece 0-9 arası rakam (tek haneli) olması gerekir.
	fmt.Println(nameControl.MatchString("kaan10")) //false

	// "emir6" ifade belirttiğimiz kurala uygun bir ifade olduğu için true.
	fmt.Println(nameControl.MatchString("emir6")) //true

	// "gokhan" ifadesi içerisinde sayı barındırmadığı için false.
	fmt.Println(nameControl.MatchString("gokhan")) //false

	// "Altan2" ifadesi büyük harf ile başladığı için false.
	fmt.Println(nameControl.MatchString("Altan2")) //false

	// "8erkay" ifadesi ise rakam sonda olması gerekirken başta olduğu için false.
	fmt.Println(nameControl.MatchString("8erkay")) //false
}
