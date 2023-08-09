package main

import (
	"fmt"
	"strconv"
)

// / Type Casting
func main() {
	var a int = 42
	b := float64(a)
	fmt.Println(b) // 42

	var c string = "42"
	d, _ := strconv.Atoi(c) // convert string to int
	fmt.Println(d)          // 42

	var e int = 42
	f := strconv.Itoa(e) // convert int to string
	fmt.Println(f)       // 42
}

// https://golangdocs.com/
