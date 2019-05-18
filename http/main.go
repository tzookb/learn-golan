package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type englisBot struct{}

type bot interface {
	getGreeting(int) string
}

func (en englisBot) getGreeting(i int) string {
	return "hello"
}

type ownWriter struct {
	content string
}

func (ow ownWriter) Write(p []byte) (n int, err error) {
	ow.content = string(p)
	return len(p), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	ow := ownWriter{}
	io.Copy(ow, resp.Body)

	fmt.Println(ow.content)
}
