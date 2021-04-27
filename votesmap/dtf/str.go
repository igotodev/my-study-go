package dtf

import (
	"bufio"
	"os"
)

// func GetString gets the file gives the array string
func GetString(fileName string) ([]string, error) {
	lines := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if fileScanner.Err() != nil {
		return nil, err
	}
	return lines, err
}
