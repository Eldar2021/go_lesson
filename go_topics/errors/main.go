package main

import (
	"errors"
	"fmt"
)

/// Hatalar (Errors)
/*
* Go dilinde hatalar, genellikle error türüyle temsil edilir.
* İşlevler, genellikle işlem sırasında bir hata oluştuğunda error değeri döndürür.
* Hataların uygun bir şekilde işlenmesi, sağlam ve güvenilir bir
* yazılım geliştirmenin önemli bir parçasıdır.
 */

func devide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("sıfıra bölme hatası")
	}
	return a / b, nil
}

func main() {
	result, error := devide(7, 0)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("Sonuç:", result)
}
