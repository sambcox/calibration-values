package calibrationvalues

import (
	"fmt"
	"strconv"
	"unicode"
)

func GetCalibrationValues(str string) int {
	firstNum := ""
	lastNum := ""
	for _, char := range str {
		if unicode.IsDigit(char) {
			num := string(char)
			if firstNum == "" {
				firstNum = num
				lastNum = num
			} else {
				lastNum = num
			}
		}
	}

	concattedNums := firstNum + lastNum

	combinedInt, err := strconv.Atoi(concattedNums)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}

	return combinedInt
}
