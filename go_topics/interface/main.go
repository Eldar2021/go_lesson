package main

import "fmt"

/// Interface
/*
* "Interface" (Arayüz), Go dilinde belirli bir türün sahip olması
* gereken metodların imza tanımlarıdır. Arayüzler, türler arasında
* birleştirme sağlar, bu sayede çok sayıda türün belirli bir davranışı
* paylaşmasına olanak tanır. Arayüzler, kodun esnekliğini artırır ve
* farklı türleri aynı şekilde işlemek için kullanılır.
*
* Go'da, bir türün belirli bir arayüzü uyguladığını ifade etmek için
* herhangi bir açıklama veya anahtar kelime kullanmanıza gerek yoktur.
* Bir tür, sadece ilgili arayüzde tanımlanan metodları uygulayarak
* otomatik olarak o arayüzü uygulamış olur.
*
* Arayüzler, genellikle kodun açık ve esnek olmasını sağlamak,
* bağımlılıkları azaltmak ve farklı türleri benzer şekillerde işlemek
* için kullanılır.
 */

type Geometry interface {
	Edge() int
}

// Pentagon defines a geometrical object
type Pentagon struct{}

// Hexagon defines a geometrical object
type Hexagon struct{}

// Octagon defines a geometrical object
type Octagon struct{}

// Decagon defines a geometrical object
type Decagon struct{}

func (p Pentagon) Edge() int { return 5 }
func (h Hexagon) Edge() int  { return 6 }
func (o Octagon) Edge() int  { return 8 }
func (d Decagon) Edge() int  { return 10 }

func Peremeter(g Geometry, i int) int {
	num := g.Edge()
	p := num * i
	return p
}

func main() {
	p := new(Pentagon)
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)

	items := [...]Geometry{p, h, o, d}

	for _, v := range items {
		fmt.Println(Peremeter(v, 5))
	}
}
