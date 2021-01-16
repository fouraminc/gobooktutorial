package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {

		//exercise 1.8
		if !strings.HasPrefix(url,"http://") {
			url = "http://" + url
			fmt.Println("No prefix")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetcg: %v\n", err)
			os.Exit(1)
		}

		//Commented line is prior to exercise 1.7
		//b, err := ioutil.ReadAll(resp.Body)

		//exercise 1.9
		fmt.Printf("%s response code for url: %s\n", resp.Status, url)
		count, err := io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
			os.Exit(1)
		}
		fmt.Printf("Written to buffer: %d\n", count)
		//Commented line is prior to exercise 1.7
		//fmt.Printf("%s", b)
	}

}
