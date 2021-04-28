// votesmap -- counts the number of identical lines (votes) in the file
// first argument file path
// for example:     $~: ./main /home/user/file.txt
// or program dir:  $~: ./main file.txt
package main

import (
	"fmt"
	"log"
	"votesmap/dtf"
	"os"
	"sort"
)

func main() {
	path := os.Args[1]
	lines, err := dtf.GetString(path)
	if err != nil {
		log.Fatal(err)
	}
	names := make(map[string]int)
	for _, line := range lines {
		/*
		_, ok := names[line]
		if !ok {
			names[line] = 1
		} else {
			names[line]++
		}
		*/
		names[line]++
	}
	value := make([]int, 0, len(names))
	for _, v := range names {
		value = append(value, v)
	}
	sort.Ints(value)
	//fmt.Println(value)
	fmt.Printf("===================================\n")
	fmt.Printf("|              VOTES🗳️             |\n")
	for i := len(value) - 1; i >= 0; i-- {
		fmt.Printf("+---------------------------------+\n")
		for name, count := range names {
			if value[i] == count {
				if value[len(value)-1] == count {
					fmt.Printf("|✅%13s | %15d|\n", name, count)
					//fmt.Printf(" ✔\n")
					//✔✖✘✗
				} else {
					fmt.Printf("|❌%13s | %15d|\n", name, count)
					//fmt.Printf(" ✗\n")
				}
			}
		}
	}
	fmt.Printf("===================================\n")
}
