package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName) // Read file, _ is used to ignore the error

	if err != nil {
		return 0, errors.New("failed to find file")
	}

	valueText := string(data) // Convert byte slice to string
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)                  // Convert float64 to string
	os.WriteFile(fileName, []byte(valueText), 0644) // Write to file
}
