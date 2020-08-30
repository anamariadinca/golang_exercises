package module1

func twoThatSum(list []int, sum int) (int, int) {
	m := make(map[int][2]int)
	key := 0
	for i := 0; i < len(list)-1; i++ {
		for j := 1; j < len(list); j++ {
			sumInList := list[i] + list[j]
			if sumInList == sum {
				m[key] = [2]int{i, j}
				key++
			}
		}
	}
	if len(m) != 0 {
		return m[0][0], m[0][1]
	}
	return -1, -1
}
