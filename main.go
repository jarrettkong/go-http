package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Printf("Error, %v\n", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
