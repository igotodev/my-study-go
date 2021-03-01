// --------------------------------------------------------------
// WebFetch - simple program which grab html page from web
// and put on created file (answer.html)
// --------------------------------------------------------------
// Run program: go run *.go [web address without prefix[https://]]
// Simple: go run *.go dev.by
// --------------------------------------------------------------
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Loadind...")
	const prefix string = "https://"
	for _, url := range os.Args[1:] {
		resp, err := http.Get(prefix + url) // get response
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body) // read the answer
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
			os.Exit(1)
		}
		f, err := os.Create("answer.html") // create a file
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		w := bufio.NewWriter(f)
		w.Write([]byte(b)) // put data into th file
		fmt.Println("Success!")
	}
}
