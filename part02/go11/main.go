package main

import "fmt"

/// Fonksiyonlar
/*
Fonksiyonlar içlerine parametre girilebilen ve işlemler yapabilen birimlerdir.
Matematikteki fonksiyonlar ile aynı mantıkta çalışan bu birimlerden bir
örneği inceleyelim.
*/

func main() {
	fmt.Println(kosh(12, 13)) // 2 + 5 sonucunu ekrana bastır
	sayName()
	fmt.Println(ishlem(6))
}

func kosh(a int, b int) int {
	return a + b // a ve b’nin toplamını döndürür.
}

func sayName() {
	fmt.Println("Eldiiar Almazbek")
}

/// Fonksiyonlar Hakkında Ayrıntılı Bilgiler
// Fonksiyon parantezi içerisine değişken tanımlanırken eğer tüm
// değişkenlerin türleri aynı ise sadece en sağdaki değişkenin tipini
// belirtmeniz yeterlidir. Örnek:

func ishlem(x int) (a, b int) {
	a = x / 2
	b = x * 2
	return
}

// Yukarıda ise isimlendirilmiş return kullandık. return tipini 
// yazdığımız paranteze bakacak olursa (x, y int) diyerek return 
// edilecek verinin fonksiyonun blokları içerisinden çekilmesini sağladık. 
// Böylece fonksiyon bloğununun sonundaki return kelimesinin yanına birşey yazmadık. 
// Bu fonksiyonumuzun çıktısı ise 5 20 olacaktır.
