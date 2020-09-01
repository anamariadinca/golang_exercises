package module2

//BubbleSort can be exprted
func BubbleSort(list []int) []int {
	var swap bool = true

	for swap == true {
		for i := 0; i < len(list); i++ {
			swap = false
			for j := 0; j < len(list)-1; j++ {
				if list[j] > list[j+1] {
					list[j], list[j+1] = list[j+1], list[j]
					swap = true
				}
			}
		}
	}
	return list
}
