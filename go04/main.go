package main

import "fmt"

/// İlişkisel Operatörler

/*
İlişkisel operatörler programlamada iki veriyi birbiriyle
karşılaştırabilmemize olanak sağlar. Karşılaştırma doğrulanıyorsa true değer,
doğrulanmıyorsa false değer alır.

== -> İki verinin eşitliği (2 == 2)
!= -> İki verinin eşitsizliği (2 != 3)
> --> Büyüktür (5 > 3)
< --> Küçüktür (4 < 6)
>= -> Büyük veya eşit (4 >= 6)
<= -> Küçük veya eşit (3 <= 8)
*/

func main() {
	fmt.Println(2 <= 3)
}
