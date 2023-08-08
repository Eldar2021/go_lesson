package main

import "fmt"

/// For Loop

// traditional for Statement
func main1() {

	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}

// for range Statement
func main2() {

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}
}

// range loop over string
func main3() {
	for range "Hello" {
		fmt.Println("Hello")
	}
}

// Infinite loop
func main() {
	main1()
	main2()
	main3()
	// i := 5
	// for {
	// 	fmt.Println("Hello")
	// 	if i == 10 {
	// 		break
	// 	}
	// 	i++
	// }
}
