// --------------------------------------------------------------
// WebFetch - simple program which grab html page from web
// and put on created file (answer.html)
// --------------------------------------------------------------
// Run program: go run *.go [web address without prefix[https://]]
// Simple: go run *.go dev.by
// --------------------------------------------------------------
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Loadind...")
	const prefix string = "https://"
	for _, url := range os.Args[1:] {
		resp, err := http.Get(prefix + url) // get response
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		f, err := os.Create("answer.html") // create a file
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		nbytes, err := io.Copy(f, resp.Body) // copy the answer
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
			os.Exit(1)
		}
		sec := time.Since(start).Seconds()
		fmt.Println("Success!")
		fmt.Printf("%s copied\n%v bytes (%fs)\n", url, nbytes, sec)
	}
}
