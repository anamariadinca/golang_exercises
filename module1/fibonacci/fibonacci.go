package module1

func getFibonacciForGivenPosition(pos int) int {
	fibonacci := []int{0, 1}
	for i := 2; i <= pos; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	return fibonacci[pos]
}
