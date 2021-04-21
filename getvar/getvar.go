// test
package getvar

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Scan user typing and convert string to float64,
// function return float64 and error
func GetFloat() (f float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	strFl, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	strFl = strings.TrimSpace(strFl)
	fl64, err := strconv.ParseFloat(strFl, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fl64, err
}

// Scan user typing and convert string to int,
// function return int and error
func GetInt() (i int, err error) {
	reader := bufio.NewReader(os.Stdin)
	strInt, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	strInt = strings.TrimSpace(strInt)
	myInt, err := strconv.ParseInt(strInt, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(myInt), err
}
