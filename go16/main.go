package main

import "fmt"

/// If-Else
/*
If ve Else kelimelerinin Türkçe karşılığına bakacak olursak;
If : Eğer, Else : Yoksa anlamına gelir. If-Else akışı koşullandırmalar
için kullanılır. Diğer dillerin aksine koşul parametresi parantezler
içine yazılmaz. Teorik kısmı bırakıp uygulama kısmına geçelim ki daha
anlaşılır olsun
*/

func main() {
	a := 5
	if a == 5 {
		fmt.Println("a'nin değeri 5'tir.")
	} else if a == 3 {
		fmt.Println("a'nin değeri 3'tir.")
	} else {
		fmt.Println("a'nin değeri 5/3 değildir.")
	}
}
