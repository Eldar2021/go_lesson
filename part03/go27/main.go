package main

import "fmt"

/// Interface
/*
"Interface", Go dilinde oldukça önemli bir kavramdır ve
Go'nun güçlü ve esnek özelliklerinden biridir. Bir interface, belirli
bir metodun davranışını tanımlayan bir türdür ve Go dilindeki tüm tiplerin
sahip olabileceği özel bir türdür.

Go'da bir interface, bir veya daha fazla metodu bildiren bir imza kümesidir.
Ancak, interfaceler kendileri metotları içermez; bunun yerine, başka bir
tipteki değerlerin bu metotları uygulamasını beklerler
*/

// Bir interface tanımlamak için şu şekilde yaparız:

type InterfaceName interface {
	Method1()
	Method2()
}

/*
Burada `InterfaceName`, tanımlanan interface'in adıdır. Interface içinde
belirtilen metotlar, başka bir tipte bu interface'i uygulayabilmek için
gerçekleştirilmesi gereken metotları temsil eder.

Daha sonra, bir tipte (struct, custom tip vb.) bu metotları uygulayarak
interface'i "impliciit" (otomatik olarak) olarak uygulayabiliriz.
Buna "interface uygulaması" denir.
*/

// Örnek olarak, bir SesCikarabilir interface'i ve bu interface'i uygulayan
// bir Kedi struct'ı düşünelim:

type SesCikarabilir interface {
	SesCikar()
}

type Kedi struct{}

func (k Kedi) SesCikar() {
	fmt.Println("Miyav!")
}

type Kopek struct{}

func (k Kopek) SesCikar() {
	fmt.Println("Vow!")
}

func main() {
	kedi := Kedi{}
	kedi.SesCikar()

	kopek := Kopek{}
	kopek.SesCikar()
}

/*
Yukarıdaki örnekte, Kopek adında başka bir struct ekledik ve bu
struct'a da SesCikar metodu ekledik. Her iki struct da SesCikarabilir
interface'ini uyguladığı için, SesCikarabilir türünde bir değişkene
atayarak hem Kedi hem de Kopek türünden nesneleri depolayabiliriz.
Bu sayede, interface sayesinde her iki türün ortak bir şekilde ele
alınabilmesini sağlarız. Bu, kodun esnekliğini artırır ve farklı türler
arasında geçiş yapmayı kolaylaştırır.

Polymorphism, Go dilindeki güçlü bir tasarım özelliğidir ve özellikle
soyutlamaya dayalı yazılım geliştirmede oldukça kullanışlıdır.
Interface'ler, Go dilindeki kodun daha modüler, yeniden kullanılabilir
ve bakımı daha kolay hale gelmesini sağlayan önemli bir yapı taşıdır.
*/
