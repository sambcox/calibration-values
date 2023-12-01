package main

import (
	"bufio"
	"fmt"
	"github.com/sambcox/calibration-values/calibrationValues"
	"os"
)

func main() {
	file, err := os.Open("calibrationinput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calibrationsum := 0
	for scanner.Scan() {
		line := scanner.Text()
		num := calibrationvalues.GetCalibrationValues(line)
		calibrationsum += num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Sum of calibrated numbers:", calibrationsum)
}
