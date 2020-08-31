package module1

import "sort"

//FactorNumber need
func FactorNumber(primes []int, num int) []int { // 2 3 5 - 30
	var refactor []int
	i := len(primes) - 1
	for i >= 0 {
		if num%primes[i] == 0 && num/primes[i] != 0 {
			refactor = append(refactor, primes[i])
			num = num / primes[i]
		} else {
			i--
		}
	}
	sort.Ints(refactor)
	return refactor
}
