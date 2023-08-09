package main

import (
	"fmt"
	"go_lesson/go_topics/packages/mymath"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := mymath.Average(xs)
	fmt.Println(avg)
}
