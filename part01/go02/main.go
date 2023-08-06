package main

import "fmt"

/*
Integer Türler
Öncelikle tüm integer türleri bir görelim;
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr

int8  -> 8-bit integers (-128 to 127)
int16 -> 16-bit integers (-32768 to 32767)
int32 -> 32-bit integers (-2147483648 to 2147483647)
int64 -> 64-bit integers (-9223372036854775808 to 9223372036854775807)

uint -> 32 or 64 bits
int -> uint (32 or 64 bits)

uint8 -> 0-255
uint16 -> 0-65535
uint32 -> 0-42967295
uint64 -> 0-18446744073709551615

Integer sayısal veri tipleri içierisinden bahsedebileceğimiz son tipler ise int ve uint.
Int ve uint veri tipleri kullanmış olduğumuz işletim sistemi 32bit ise 32bit değer
alırlar, 64bit ise 64bit değer alırlar. Sayısal bir değer atanacağı zaman en çok
kullanılan veri tipleridir. Genellikle int daha çok kullanılır. Eğer çok meşakkatli
bir program yazmayacaksanız int kullanmanız önerilir.

Byte Veri Tipi: uint8 ile aynıdır.
Rune: int32 ile aynıdır. Unicode karakter kodlarını ifade eder.
*/

/*
Float türleri integer türlerden farklı olarak küsüratlı sayıları tutar. Örnek: 3.14
float32: 32bitlik değer alabilir.
float64: 64 değer alabilir.
*/
/*
Complex Türler
Complex türleri içerisinde gerçel küsüratlı (float) ve sanal sayılar barındırabilir.
Türkçe’de karmaşık sayılar diye adlandırılır.

complex64: Gerçel float32 ve sanal sayı değeri barındırır.
complex128: Gerçel float64 ve sanal sayı değeri barındırır.
Sayısal türler bu şekildedir.
*/
/*
BOOLEAN VERİ TİPİ
Boolean yani mantıksal veri tipi bir durumun var olması halinde olumlu (true) değer,
var olmaması halinde olumsuz (false) değer alan veri tipidir.
*/
/*
STRING VERİ TİPİ
String yani dizgi veri tipi içerisinde metinsel ifadeler barındırır.
Örnek olarak “Golang çok güzel ama ingilicce”. String veri tipi değeri çift tırnak ( “Değer” ) içine yazılır.
Diğer dillerdeki gibi tek tırnak ( ‘Değer’ ) insiyatifi yoktur. Tek tırnakla kullanım başka bir amaç içindir.
İlerde onu da göstereceğim.
*/
/*
Veri tipleri atanacak değerlerimizi RAM üzerinde depolamak için kullandığımız araçlardır.
Tam sayı değerler için Integer veri tiplerini, ondalık sayılar için Float veri tiplerini,
mantıksal değerler için Boolean veri tipini, metinsel değerler için String veri tipini kullanırız.
Karmaşık sayı değerleri için ise Complex veri tipini kullanırız.
*/

func main() {
	// (-128 to 127)
	var intVar int8 = 125
	fmt.Println("int8 ", intVar)

	// float32: 32bitlik değer alabilir.
	var floatVar float32 = 12.3
	fmt.Println("float32 ", floatVar)

	// "{value}"
	var strVar string = "Eldiiar"
	fmt.Println("string ", strVar)

	// bool true-false
	var boolVar bool = true
	fmt.Println("bool ", boolVar)
}
