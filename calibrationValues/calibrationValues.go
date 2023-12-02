package calibrationvalues

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetCalibrationValues(str string) int {
	forwardPattern := `one|two|three|four|five|six|seven|eight|nine|[1-9]`
	backwardPattern := `enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|[1-9]`

	reversedStr := reverseString(str)

	forwardRegex := regexp.MustCompile(forwardPattern)
	backwardRegex := regexp.MustCompile(backwardPattern)

	matches := forwardRegex.FindAllStringSubmatch(str, -1)
	backwardMatches := backwardRegex.FindAllStringSubmatch(reversedStr, -1)

	if reverseString(backwardMatches[0][0]) != matches[len(matches)-1][0] {
		matches = append(matches, []string{reverseString(backwardMatches[0][0])})
	}

	if len(matches) == 0 {
		return 0
	}

	firstNum := matches[0][0]
	lastNum := matches[len(matches)-1][0]
	numReference := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if _, ok := numReference[firstNum]; ok {
		firstNum = numReference[firstNum]
	}

	if _, ok := numReference[lastNum]; ok {
		lastNum = numReference[lastNum]
	}

	concattedNums := firstNum + lastNum

	combinedInt, err := strconv.Atoi(concattedNums)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return 0
	}

	return combinedInt
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
