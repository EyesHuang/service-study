package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "listen port")
	flag.Parse()
	if err := run(*port); err != nil {
		fmt.Println("error to run service")
	}
}

func run(port int) error {
	srv := NewServer()
	return http.ListenAndServe(fmt.Sprintf(":%d", port), srv)
}
