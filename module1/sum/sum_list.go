package module1

func ComputeSumOfList(list []int) int {
	sum := 0

	for _, n := range list {
		sum += n
	}

	return sum
}

func ComputeSumRecursive(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + ComputeSumRecursive(list[1:])
}
