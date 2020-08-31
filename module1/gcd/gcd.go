package module1

import (
	factoring "golang_exercises/module1/factoring"
	"math"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139}

func gcdUsingFactoring(first int, second int) int {
	mapFirst := countEachPrimeNr(factoring.FactorNumber(primes, first))
	mapSecond := countEachPrimeNr(factoring.FactorNumber(primes, second))
	cod := 1
	for _, prime := range primes {
		if mapFirst[prime] != 0 && mapSecond[prime] != 0 {
			exponent := math.Min(float64(mapFirst[prime]), float64(mapSecond[prime]))
			cod *= int(math.Pow(float64(prime), exponent))
		}
	}
	return cod
}

func countEachPrimeNr(decomposedNumber []int) map[int]int {
	countMap := make(map[int]int)
	for _, prime := range decomposedNumber {
		countMap[prime]++
	}
	return countMap
}

//GcdUsingEuclidean can be used in main.go
func GcdUsingEuclidean(first int, second int) int {
	for second > 0 {
		first, second = second, first%second
	}
	return first
}
