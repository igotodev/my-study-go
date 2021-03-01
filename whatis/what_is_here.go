// what_is_here -- is simple microprogram 
// show files in selected directory
package main

import (
	"fmt"
	"io/ioutil"
)

func getFiles(path string) ([]string, error) {
	var files []string
	dirFiles, err := ioutil.ReadDir(fmt.Sprintf("%s", path))
	if err != nil {
		return nil, err
	}

	for _, i := range dirFiles {
		files = append(files, i.Name())
	}
	return files, nil
}

func main() {

	var your_path string
	fmt.Println("Please, enter directory")	
	fmt.Println("\"Simple: /home/user/Downloads\"")
	fmt.Printf("Path to directory: ")
	fmt.Scanf("%s", &your_path)
	f, _ := getFiles(your_path)
	for i := 0; i < len(f); i++ {
		fmt.Printf("%d) ", i+1)
		fmt.Printf(f[i] + "\n")
	}
	fmt.Printf("\"%s\" has %d file(s)/directory(es)\n",
		your_path, len(f))
}
