package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom Writer

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// bs := []byte{} - won't work as it'll create a zero-size buffer
	/*
		SizeOfBuffer := 9 * 1024 * 1024
		bs := make([]byte, SizeOfBuffer)
		n, err := resp.Body.Read(bs)
		fmt.Println(n, err)
		fmt.Println(string(bs))
	*/
	// io.Copy(os.Stdout, resp.Body)
	// custom write
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	println(string(bs))
	return len(bs), nil
}
