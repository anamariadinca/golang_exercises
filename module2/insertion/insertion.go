package module2

//Insertion can be exported
func Insertion(list []int) []int {
	var sortedList = []int{list[0]}
	for i := 1; i < len(list); i++ {
		sortedList = insertInSortedList(sortedList, list[i])
	}
	return sortedList
}

//InsertionTutorial can be exported
func InsertionTutorial(list []int) []int {
	var sortedList = []int{list[0]}
	for i := 1; i < len(list); i++ {
		sortedList = insertInSortedListTutorial(sortedList, list[i])
	}
	return sortedList
}

func insertInSortedList(sortedList []int, inNum int) []int {
	inserted := false
	for i := 0; i < len(sortedList); i++ {
		if inNum <= sortedList[i] {
			aux := sortedList[len(sortedList)-1]
			for j := len(sortedList) - 1; j > i; j-- {
				sortedList[j] = sortedList[j-1]
			}
			sortedList[i] = inNum
			sortedList = append(sortedList, aux)
			inserted = true
			break
		}
	}
	if !inserted {
		sortedList = append(sortedList, inNum)
	}
	return sortedList
}

func insertInSortedListTutorial(sortedList []int, inNum int) []int {
	inserted := false
	for i := 0; i < len(sortedList); i++ {
		if inNum <= sortedList[i] {
			sortedList = append(sortedList[:i], append([]int{inNum}, sortedList[:i]...)...)
		}
	}
	if !inserted {
		sortedList = append(sortedList, inNum)
	}
	return sortedList
}
