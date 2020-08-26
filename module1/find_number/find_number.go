package module1

func NumInList(numList []int, num int) bool {
	for _, n := range numList {
		if n == num {
			return true
		} 
	} 
	return false
}