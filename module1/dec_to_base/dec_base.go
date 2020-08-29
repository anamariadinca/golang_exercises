package module1

import "strconv"

func decToAnyBase(dec int, base int) string {
	convertedNumber := ""

	for dec > 0 {
		convertedNumber = ConvertReminderToString(dec%base) + convertedNumber
		dec = dec / base
	}

	return convertedNumber
}

//ConvertReminderToString : In case the conversion is for a base bigger than 10, this returns the letter corresponding for the reminder
//eg: there is a reminder of 12 in case of a conversion in base 16, this will return C instead of 12
func ConvertReminderToString(reminder int) string {
	switch reminder {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return strconv.Itoa(reminder)
	}
}
