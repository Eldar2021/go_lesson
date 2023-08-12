package main

import "fmt"

/// Testing
/*
* Hücrelerin vücumudaki yapı birimi olduğu gibi, aynı şekilde her bileşen
* de yazılımın birer parçasıdır. Yazılımın sağlıklı bir şekilde çalışabilmesi için, her
* bileşenin güvenilir bir şekilde çalışması gerekir. Aynı şekilde vücudumuzun sağlığı
* hücrelerin güvenilirliği ve verimliliğine bağlı olduğu gibi, yazılımın düzgün
* çalışması bileşenlerin güvenilirliği ve verimliliğine bağlıdır.
* Biraz biyoloji dersi gibi oldu ama sonuçta aynı mantığı yürütebiliriz.
 */

func sayHello(name string, lastname string) string {
	fullname := name + " " + lastname
	return fullname
}

func main() {
	result := sayHello("Eldiiar", "Almazbek")
	fmt.Println(result)
}
