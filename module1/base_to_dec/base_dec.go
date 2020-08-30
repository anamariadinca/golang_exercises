package module1

import (
	"math"
	"strconv"
)

func baseToDeck(baseNum string, base int) int {
	decNum := 0
	power := len(baseNum) - 1
	for i := 0; i < len(baseNum); i++ {
		basePower := math.Pow(float64(base), float64(power))
		multiplier, err := ConvertMultiplierToInt(baseNum[i])
		if err == nil {
			decNum += int(basePower) * multiplier
		}
		power--
	}
	return decNum
}

// ConvertMultiplierToInt converts the
func ConvertMultiplierToInt(numB byte) (int, error) {
	numS := string(numB)
	switch numS {
	case "A":
		return 10, nil
	case "B":
		return 11, nil
	case "C":
		return 12, nil
	case "D":
		return 13, nil
	case "E":
		return 14, nil
	case "F":
		return 15, nil
	default:
		return strconv.Atoi(numS)
	}
}
