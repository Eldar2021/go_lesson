package main

import "fmt"

/// Panic / Recover
/*
* `panic`, beklenmedik bir durumda veya programın doğru bir şekilde devam
* edemeyeceği bir noktada kullanılan bir mekanizmadır. panic çağrısı
* yapıldığında, normal program akışı kesilir ve ardından "panik durumu"
* ortaya çıkar. Ancak, panic durumunun iyi yönetilmemesi programın çökmesine
* neden olabilir.
 */
/*
* `recover`, `panic` durumlarını kontrol etmek ve programın çalışmasını
* sürdürmek için kullanılır. `recover` fonksiyonu, bir `defer` ifadesi içinde
* kullanıldığında, bir `panik` durumunda programın devam etmesine izin verir
* ve `panik` durumunun nedenini yakalayabilir.
 */

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panik durumu:", r)
		}
	}()

	fmt.Println("Program başladı")

	panic("bir panik durumu oluştu")
}

/*
* Bu konularda dikkatli olmak, hatasız ve güvenilir kod yazmak için önemlidir.
* Hataları uygun bir şekilde yönetmek ve panik durumlarını kontrol altına almak,
* programınızın sağlam ve hata toleranslı olmasını sağlar.
 */
