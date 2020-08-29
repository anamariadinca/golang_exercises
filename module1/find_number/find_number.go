package module1

func numInList(numList []int, num int) bool {
	for _, n := range numList {
		if n == num {
			return true
		}
	}
	return false
}
