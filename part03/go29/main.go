package main

import "fmt"

/// Generics
/*
Go dilinde generics, belirli bir tip yerine herhangi bir tip için kod
yazmanıza olanak tanıyan bir özelliktir. Bu, kodun daha genel ve yeniden
kullanılabilir hale gelmesini sağlar.

Go'da generics, T gibi bir genel değişken kullanılarak tanımlanır.
T yerine herhangi bir tip kullanılabilir. Örneğin, aşağıdaki kod,
T için herhangi bir tipteki bir diziyi sıralamak için bir işlev tanımlar:
*/

func SortArray[T comparable](arr []T) {
	fmt.Println(arr)
}

/*
Bu işlevi, int, string veya float64 gibi herhangi bir tipteki bir diziyi
sıralamak için kullanabilirsiniz. Örneğin, aşağıdaki kod, bir int dizisini
sıralamak için işlevi çağırır:
*/

func main() {
	arr := []int{1, 2, 4, 5}
	SortArray(arr) // [1 2 4 5]
}

/*
Go'da generics, kodunuzu daha genel ve yeniden kullanılabilir hale getirmek
için harika bir yoldur. Ancak, generics'in performans üzerindeki etkilerini
dikkate almak önemlidir. Genel olarak, generics'in performans üzerinde önemli
bir etkisi yoktur, ancak bazı durumlarda performansı etkileyebilir.
*/
