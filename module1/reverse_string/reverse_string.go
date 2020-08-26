package module1

func reverseString(word string) string {
	var result string
	for _, l := range word {
		result = string(l) + result
	}
	return result
}
