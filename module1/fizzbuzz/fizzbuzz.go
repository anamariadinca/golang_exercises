package module1

import "strconv"

func fizzBuzz(number int) string {
	var s string
	for i := 1; i < number+1; i++ {
		rem3 := i % 3
		rem5 := i % 5
		if rem3 == 0 || rem5 == 0 {
			if rem3 == 0 {
				s += "Fizz"
			}
			if rem5 == 0 {
				s += "Buzz"
			}
		} else {
			s += strconv.Itoa(i)
		}
		if i < number {
			s += ", "
		}
	}
	return s + "\n"
}
