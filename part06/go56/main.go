package main

import "go_lesson/part06/go56/mypackage"

/// Paket (Kütüphane) Yazmak
/*
* Bu bölümde Go üzerinde nasıl kendi paketimizi (kütüphanemizi)
* oluşturacağımıza bakacağız.
*
* Bir Paketin Özellikleri:
* - İçerisinde .go dosyaları bulunan bir klasördür.
* - Diğer projeler tarafından içe aktarılabilir.
* - Dışa aktarılabilen veya aktarılamayan veriler içerir.
* - Açık kaynaktır.
* - main() fonksiyonu içermez.
* - package main değildir.
 */

/*
* Paket oluştururken dikkat etmemiz gereken prensipler vardır.
* Bunlar şık ve basit kod yazımı, dışarıdan kullanımı basit, mümkün
* olduğunca diğer paketlere bağımsız olmasıdır. Bu prensiplere dikkat ederek daha
* iyi bir paket yazabilirsiniz.
*
* `mypackage`
 */

func main() {
	mypackage.SayHello()
}

/*
* Git Sisteminde Kütüphane Paylaşımı
*
* Oluşturduğumuz kütüphaneyi Github, Gitlab, Bitbucket vb. sitelerde barındırarak
* diğer geliştiricilerinde kütüphanelerinizden faydalanmasını sağlayabilirsiniz.
*
* Bunun için kütüphanenizin isminde bir repo oluşturup, içerisinde
* Go dosyalarınızı yükleyin. Daha sonra go get github.com/id/repoismi
* şeklinde projenize import edebilirsiniz.
 */
