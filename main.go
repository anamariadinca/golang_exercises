package main

import (
	"fmt"
	gdc "golang_exercises/module1/gcd"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		fmt.Printf("%d\n", gdc.GcdUsingEuclidean(a, b))
	}
}
